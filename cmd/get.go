package cmd

import (
	"github.com/spf13/cobra"
	"moaictl/cmd/accl"
	"moaictl/cmd/computenode"
	"moaictl/cmd/flavor"
	"moaictl/cmd/mafversion"
	"moaictl/cmd/schedulerinfo"
	"moaictl/cmd/schedulinggroup"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get resources",
}

func init() {
	getCmd.AddCommand(accl.GetAcclCmd)
	getCmd.AddCommand(computenode.GetCpnCmd)
	getCmd.AddCommand(flavor.GetFlavorCmd)
	getCmd.AddCommand(mafversion.GetMafVersionCmd)
	getCmd.AddCommand(schedulinggroup.GetSchedulingGroupCmd)
	getCmd.AddCommand(schedulerinfo.GetSchedulerInfoCmd)
}
