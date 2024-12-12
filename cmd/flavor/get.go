package flavor

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetFlavorCmd represents the "get accl history" command
var GetFlavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "get a flavor",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			flavorName = args[0]
		}

		if err := getFlavor(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	GetFlavorCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed flavor info (optional)")
}

func getFlavor() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionFlavor
	if flavorName != "" {
		url += URLSeperator + flavorName
	}

	resp, err := client.RequestGet(url)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	if flavorName != "" {
		return PrintFlavorK8sStyle(resp.Body)
	}

	return PrintFlavorListK8sStyle(resp.Body)
}
