package devicefailurehistory

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
	"net/http"
)

// GetDeviceFailureHistoryCmd represents the "get accl history" command
var GetDeviceFailureHistoryCmd = &cobra.Command{
	Use:   "device-failure-history",
	Short: "get an device failure history",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getDeviceFailureHistory(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

var (
	sgName     string
	isDetailed bool
)

func init() {
	GetDeviceFailureHistoryCmd.Flags().StringVarP(&sgName, "scheduling-group-name", "s", "", "name of the scheduling group (optional)")
	GetDeviceFailureHistoryCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed accelerator infos (optional)")
}

func getDeviceFailureHistory() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionComputeNode + URLPartitionDeviceFailureHistory
	if sgName != "" {
		url += URLQueryParamSchedulingGroupName + sgName
	}

	resp, err := client.RequestDo(http.MethodGet, url, nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	return PrintDeviceFailureHistoryListK8sStyle(resp.Body)
}
