package accl

import (
	"fmt"
	"github.com/spf13/cobra"
	client "moaictl/pkg/common/client"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// DeleteAcclCmd represents the "delete accl" command
var DeleteAcclCmd = &cobra.Command{
	Use:   "accl",
	Short: "delete an accl resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			acclName = args[0]
		}

		if err := deleteAccl(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Accelerator \"%s\" deleted\n", acclName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
}

func deleteAccl() error {
	if acclName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionAccelerator + URLPartitionName + URLSeperator + acclName

	resp, err := client.RequestDo(http.MethodDelete, url, nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	return nil
}
