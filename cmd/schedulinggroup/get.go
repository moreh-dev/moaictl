package schedulinggroup

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
	"net/http"
)

// GetSchedulingGroupCmd represents the "get scheduling-group" command
var GetSchedulingGroupCmd = &cobra.Command{
	Use:   "schd-group",
	Short: "get a scheduling group",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			sgName = args[0]
		}

		if err := getSchedulingGroup(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	GetSchedulingGroupCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed scheduling group info (optional)")
}

func getSchedulingGroup() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionSchedulingGroup
	if sgName != "" {
		url += URLPartitionName + URLSeperator + sgName
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

	if sgName != "" {
		return PrintSchedulingGroupK8sStyle(resp.Body)
	}

	return PrintSchedulingGroupListK8sStyle(resp.Body)
}
