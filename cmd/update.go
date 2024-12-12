package cmd

import (
	"github.com/spf13/cobra"
	"moaictl/cmd/flavor"
	mafversion "moaictl/cmd/mafversion"
	sg "moaictl/cmd/schedulinggroup"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a resource",
}

func init() {
	updateCmd.AddCommand(flavor.UpdateFlavorCmd)
	updateCmd.AddCommand(mafversion.UpdateMafVersionCmd)
	updateCmd.AddCommand(sg.UpdateSchedulingGroupCmd)
}
