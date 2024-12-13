package mafversion

import (
	"fmt"
	"github.com/spf13/cobra"
	client "moaictl/pkg/common/client"
	"net/http"

	. "moaictl/pkg/common/config"
	. "moaictl/pkg/common/constants"
)

// DeleteMafVersionCmd represents the "delete mafVersion" command
var DeleteMafVersionCmd = &cobra.Command{
	Use:   "maf-version",
	Short: "delete an mafVersion resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tag = args[0]
		}

		if err := deleteMafVersion(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("MafVersion \"%s\" deleted\n", tag)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
}

func deleteMafVersion() error {
	if tag == "" {
		return fmt.Errorf("resource(s) were provided, but no name was specified")
	}

	url := Config.APIServerAddress + URLVersionV1 + URLPartitionMafVersion + URLSeperator + tag

	resp, err := client.RequestDo(http.MethodDelete, url, nil)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer client.CloseResponseBody(resp)

	return nil
}
