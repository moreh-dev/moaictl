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

// AddMafVersionCmd represents the "add mafVersion" command
var AddMafVersionCmd = &cobra.Command{
	Use:   "maf-version",
	Short: "add an maf version resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tag = args[0]
		}

		if err := addMafVersion(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nMafVersion \"%s\" added\n", tag)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	// Add flags specific to "add mafVersion"
	AddMafVersionCmd.Flags().BoolVarP(&enabled, "enabled", "e", false, "enable/disable the maf version resource (optional)")
	AddMafVersionCmd.Flags().BoolVarP(&latest, "latest", "l", false, "latest the maf version resource (optional)")
	// TODO: MafEnvs
}

func addMafVersion() error {
	if tag == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionMafVersion

	payload := utils.MafVersion{
		Tag:     tag,
		Enabled: enabled,
		Latest:  latest,
		// TODO: MafEnvs
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := client.RequestDo(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	return utils.PrintMafVersionK8sStyle(resp.Body)
}
