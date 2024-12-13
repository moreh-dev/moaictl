package schedulinggroup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	"moaictl/pkg/common/utils"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// AddSchedulingGroupCmd represents the "add schedulingGroup" command
var AddSchedulingGroupCmd = &cobra.Command{
	Use:   "schd-group",
	Short: "add an scheduling group resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			sgName = args[0]
		}

		if err := addSchedulingGroup(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nSchedulingGroup \"%s\" added\n", sgName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	// Add flags specific to "add scheduling-group"
	AddSchedulingGroupCmd.Flags().StringVarP(&schedulingPolicy, "scheduling-policy", "s", "default", "scheduling policy (optional)")
	// TODO: fix default OptimizedParallel to Default. fix mam too.
	AddSchedulingGroupCmd.Flags().StringVarP(&allocationPolicy, "allocation-policy", "a", "OptimizedParallel", "allocation policy (optional)")
}

func addSchedulingGroup() error {
	if sgName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionSchedulingGroup

	payload := utils.SchedulingGroup{
		Name:             sgName,
		SchedulingPolicy: schedulingPolicy,
		AllocationPolicy: allocationPolicy,
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

	return utils.PrintSchedulingGroupK8sStyle(resp.Body)
}
