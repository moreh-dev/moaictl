package nodefailurehistory

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetNodeFailureHistoryCmd represents the "get accl history" command
var GetNodeFailureHistoryCmd = &cobra.Command{
	Use:   "node-failure-history",
	Short: "get an node failure history",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getNodeFailureHistory(); err != nil {
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
	GetNodeFailureHistoryCmd.Flags().StringVarP(&sgName, "scheduling-group-name", "s", "", "name of the scheduling group (optional)")
	GetNodeFailureHistoryCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed accelerator infos (optional)")
}

func getNodeFailureHistory() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionComputeNode + URLPartitionNodeFailureHistory
	if sgName != "" {
		url += URLQueryParamSchedulingGroupName + sgName
	}

	resp, err := client.RequestGet(url)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	return PrintNodeFailureHistoryListK8sStyle(resp.Body)
}
