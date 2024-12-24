package history

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
	"net/http"
)

// GetHistoryCmd represents the "get accl history" command
var GetHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "get an accl history",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getHistory(); err != nil {
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
	GetHistoryCmd.Flags().StringVarP(&sgName, "scheduling-group-name", "s", "", "name of the scheduling group (optional)")
	GetHistoryCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed accelerator infos (optional)")
}

func getHistory() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionAccelerator + URLPartitionHistory
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

	return PrintAcclHistoryListK8sStyle(resp.Body)
}
