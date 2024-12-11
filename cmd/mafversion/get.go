package mafversion

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetMafVersionCmd represents the "get accl history" command
var GetMafVersionCmd = &cobra.Command{
	Use:   "maf-version",
	Short: "get a maf version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tag = args[0]
		}

		if err := getMafVersion(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	GetMafVersionCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed maf version info (optional)")
}

func getMafVersion() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionMafVersion
	if tag != "" {
		url += URLSeperator + tag
	}

	resp, err := client.RequestGet(url)
	if err != nil {
		return err
	}
	defer client.CloseResponseBody(resp)

	if isDetailed {
		return PrettyPrintJSON(resp.Body)
	}

	if tag != "" {
		return PrintMafVersionK8sStyle(resp.Body)
	}

	return PrintMafVersionListK8sStyle(resp.Body)
}
