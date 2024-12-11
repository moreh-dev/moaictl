package flavor

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

// UpdateFlavorCmd represents the "update flavor" command
var UpdateFlavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "update an flavor resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			flavorName = args[0]
		}

		if err := updateFlavor(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nFlavor \"%s\" updated\n", flavorName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	// Update flags specific to "update flavor"
	UpdateFlavorCmd.Flags().IntVarP(&deviceCount, "device-count", "c", 0, "device count of the flavor (required)")
	UpdateFlavorCmd.Flags().StringVarP(&schedulingGroupName, "scheduling-group-name", "s", "default", "name of the scheduling group (optional)")
	// TODO: MAFEnvs

	if err := UpdateFlavorCmd.MarkFlagRequired("device-count"); err != nil {
		panic(err)
	}
}

func updateFlavor() error {
	if flavorName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionFlavor

	payload := utils.Flavor{
		Name:                flavorName,
		SchedulingGroupName: schedulingGroupName,
		DeviceCount:         deviceCount,
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
	// return utils.PrintFlavorK8sStyle(resp.Body)
	return nil
}
