package schedulerinfo

import (
	"fmt"
	"github.com/spf13/cobra"

	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetSchedulerInfoCmd represents the "get scheduler-info" command
var GetSchedulerInfoCmd = &cobra.Command{
	Use:   "schd-info",
	Short: "get a scheduler info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			siName = args[0]
		}

		if err := getSchedulerInfo(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

var (
	siName     string
	isDetailed bool
)

func init() {
	GetSchedulerInfoCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed scheduler info info (optional)")
}

func getSchedulerInfo() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionSchedulerInfo
	if siName != "" {
		url += URLSeperator + siName
	}

	resp, err := client.RequestGet(url)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	if siName != "" {
		return PrintSchedulerInfoK8sStyle(resp.Body)
	}

	return PrintSchedulerInfoListK8sStyle(resp.Body)
}
