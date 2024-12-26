package computenode

import (
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/cmd/computenode/devicefailurehistory"
	"moaictl/cmd/computenode/nodefailurehistory"
	"moaictl/pkg/common/client"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	. "moaictl/pkg/common/utils"
)

// GetCpnCmd represents the "get cpn" command
var GetCpnCmd = &cobra.Command{
	Use:   "cpn",
	Short: "get a compute node resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cpnName = args[0]
		}

		if err := getCpn(); err != nil {
			fmt.Println(err)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	GetCpnCmd.AddCommand(nodefailurehistory.GetNodeFailureHistoryCmd)
	GetCpnCmd.AddCommand(devicefailurehistory.GetDeviceFailureHistoryCmd)

	// Add flags specific to "get cpn"
	GetCpnCmd.Flags().BoolVarP(&isDetailed, "detail", "d", false, "get detailed accelerator infos (optional)")
}

func getCpn() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionComputeNode
	if cpnName != "" {
		url += URLPartitionName + URLSeperator + cpnName
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

	if cpnName != "" {
		return PrintCpnK8sStyle(resp.Body)
	}

	return PrintCpnListK8sStyle(resp.Body)
}
