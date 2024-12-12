package cmd

import (
	"github.com/spf13/cobra"
	"moaictl/cmd/accl"
	"moaictl/cmd/flavor"
	"moaictl/cmd/mafversion"
	"moaictl/cmd/schedulinggroup"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a resource",
}

func init() {
	addCmd.AddCommand(accl.AddAcclCmd)
	addCmd.AddCommand(flavor.AddFlavorCmd)
	addCmd.AddCommand(mafversion.AddMafVersionCmd)
	addCmd.AddCommand(schedulinggroup.AddSchedulingGroupCmd)
}
