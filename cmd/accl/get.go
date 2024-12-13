package accl

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/cmd/accl/history"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetAcclCmd represents the "get accl" command
var GetAcclCmd = &cobra.Command{
	Use:   "accl",
	Short: "get an accl resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			acclName = args[0]
		}

		if err := getAccl(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	GetAcclCmd.AddCommand(history.GetHistoryCmd)

	// Add flags specific to "get accl"
	GetAcclCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed accelerator infos (optional)")
}

func getAccl() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionAccelerator
	if acclName != "" {
		url += URLPartitionName + URLSeperator + acclName
	}

	resp, err := client.RequestGet(url)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	if acclName != "" {
		return PrintAcclK8sStyle(resp.Body)
	}

	return PrintAcclListK8sStyle(resp.Body)
}
