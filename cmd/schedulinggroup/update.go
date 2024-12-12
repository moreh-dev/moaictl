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

// UpdateSchedulingGroupCmd represents the "update schedulingGroup" command
var UpdateSchedulingGroupCmd = &cobra.Command{
	Use:   "schd-group",
	Short: "update an scheduling-group resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			sgName = args[0]
		}

		if err := updateSchedulingGroup(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\nSchedulingGroup \"%s\" updated\n", sgName)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	// Update flags specific to "update scheduling group"
	UpdateSchedulingGroupCmd.Flags().StringVarP(&schedulingPolicy, "scheduling-policy", "s", "default", "scheduling policy (optional)")
	UpdateSchedulingGroupCmd.Flags().StringVarP(&allocationPolicy, "allocation-policy", "a", "OptimizedParallel", "allocation policy (optional)")
}

func updateSchedulingGroup() error {
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

	resp, err := client.RequestDo(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	// TODO: fix mam, update must return response body
	//return utils.PrintSchedulingGroupK8sStyle(resp.Body)
	return nil
}
