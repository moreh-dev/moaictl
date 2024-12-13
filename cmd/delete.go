package cmd

import (
	"github.com/spf13/cobra"
	"moaictl/cmd/accl"
	"moaictl/cmd/flavor"
	"moaictl/cmd/mafversion"
	"moaictl/cmd/schedulinggroup"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a resource",
}

func init() {
	deleteCmd.AddCommand(accl.DeleteAcclCmd)
	deleteCmd.AddCommand(flavor.DeleteFlavorCmd)
	deleteCmd.AddCommand(mafversion.DeleteMafVersionCmd)
	deleteCmd.AddCommand(schedulinggroup.DeleteSchedulingGroupCmd)

}
