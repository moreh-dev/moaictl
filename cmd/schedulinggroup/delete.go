package schedulinggroup

import (
	"fmt"
	"github.com/spf13/cobra"
	client "moaictl/pkg/common/client"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// DeleteSchedulingGroupCmd represents the "delete schedulingGroup" command
var DeleteSchedulingGroupCmd = &cobra.Command{
	Use:   "schd-group",
	Short: "delete an schedulingGroup resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			sgName = args[0]
		}

		if err := deleteSchedulingGroup(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("SchedulingGroup \"%s\" deleted\n", sgName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
}

func deleteSchedulingGroup() error {
	if sgName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionSchedulingGroup + URLPartitionName + URLSeperator + sgName

	resp, err := client.RequestDo(http.MethodDelete, url, nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	return nil
}
