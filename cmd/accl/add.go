package accl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"

	"moaictl/pkg/common/client"
	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
	"moaictl/pkg/common/utils"
)

var (
	namespace string
)

// AddAcclCmd represents the "create accl" command
var AddAcclCmd = &cobra.Command{
	Use:   "accl",
	Short: "Add an accl resource",
	Long:  `Add an accl resource. Either flavor or device count must be provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := addAccl(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nAccelerator added\n")
		}
	},
}

func init() {
	// Add flags specific to "add accl"
	AddAcclCmd.Flags().StringVarP(&mafVersion, "maf-version", "v", "", "Version of moreh ai framework (required)")
	AddAcclCmd.Flags().IntVarP(&priority, "priority", "p", 0, "Priority of the accl resource (required)")
	AddAcclCmd.Flags().StringVarP(&flavorName, "flavor-name", "f", "", "Name of the accl flavor (optional)")
	AddAcclCmd.Flags().StringVarP(&schedulingGroupName, "scheduling-group-name", "s", "default", "Name of the scheduling group (optional)")
	AddAcclCmd.Flags().IntVarP(&gpus, "device-count", "d", 0, "device count(optional)")
	// TODO: MafEnvs

	if err := AddAcclCmd.MarkFlagRequired("maf-version"); err != nil {
		panic(err)
	}
	if err := AddAcclCmd.MarkFlagRequired("priority"); err != nil {
		panic(err)
	}
}

func addAccl() error {
	url := Config.APIServerAddress + URLVersionV1 + URLPartitionAccelerator

	payload := utils.AcclItem{
		FlavorName:          flavorName,
		MafVersion:          mafVersion,
		SchedulingGroupName: schedulingGroupName,
		Priority:            priority,
		GPUs:                gpus,
		// TODO: MafEnvs
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

	return utils.PrintAcclK8sStyle(resp.Body)
}
