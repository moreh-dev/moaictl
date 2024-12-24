package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"moaictl/pkg/common/client"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to the moai acceleartor manager api server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := login(); err != nil {
			fmt.Println(err)
		}
	},
}

func login() error {
	fmt.Print("Enter your MoAI token: ")
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	err = client.SaveToken(token)
	if err != nil {
		return err
	}

	return nil
}
