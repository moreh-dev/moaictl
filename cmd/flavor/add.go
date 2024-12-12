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

// AddFlavorCmd represents the "add flavor" command
var AddFlavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "add an flavor resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			flavorName = args[0]
		}

		if err := addFlavor(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nFlavor \"%s\" added\n", flavorName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	// Add flags specific to "add flavor"
	AddFlavorCmd.Flags().IntVarP(&deviceCount, "device-count", "c", 0, "device count of the flavor (required)")
	AddFlavorCmd.Flags().StringVarP(&schedulingGroupName, "scheduling-group-name", "s", "default", "name of the scheduling group (optional)")
	// TODO: MAFEnvs
	// TODO: yaml 혹은 json으로 일괄적으로 받아서 적용하는 방식 고려

	if err := AddFlavorCmd.MarkFlagRequired("device-count"); err != nil {
		panic(err)
	}
}

func addFlavor() error {
	if flavorName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionFlavor

	payload := utils.Flavor{
		Name:                flavorName,
		SchedulingGroupName: schedulingGroupName,
		DeviceCount:         deviceCount,
		// TODO: MAFEnvs
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

	return utils.PrintFlavorK8sStyle(resp.Body)
}
