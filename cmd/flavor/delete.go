package flavor

import (
	"fmt"
	"github.com/spf13/cobra"
	client "moaictl/pkg/common/client"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// DeleteFlavorCmd represents the "delete flavor" command
var DeleteFlavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "delete an flavor resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			flavorName = args[0]
		}

		if err := deleteFlavor(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Flavor \"%s\" deleted\n", flavorName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
}

func deleteFlavor() error {
	if flavorName == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionFlavor + URLSeperator + flavorName

	resp, err := client.RequestDo(http.MethodDelete, url, nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	return nil
}
