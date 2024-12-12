package mafversion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	"moaictl/pkg/common/utils"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// UpdateMafVersionCmd represents the "update mafVersion" command
var UpdateMafVersionCmd = &cobra.Command{
	Use:   "maf-version",
	Short: "update an maf-version resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tag = args[0]
		}

		if err := updateMafVersion(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nMafVersion \"%s\" updated\n", tag)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	// Update flags specific to "update mafVersion"
	UpdateMafVersionCmd.Flags().BoolVarP(&enabled, "enabled", "e", false, "enable/disable the maf version resource (optional)")
	UpdateMafVersionCmd.Flags().BoolVarP(&latest, "latest", "l", false, "latest the maf version resource (optional)")
	// TODO: MAFEnvs
}

func updateMafVersion() error {
	if tag == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionMafVersion

	payload := utils.MafVersion{
		Tag:     tag,
		Enabled: enabled,
		// TODO: fix mam, updating latest to true must also set the previous latest version to false
		Latest: latest,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := client.RequestDo(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	// TODO: fix mam, update must return response body
	//return utils.PrintMafVersionK8sStyle(resp.Body)
	return nil
}
