package main

import (
	"github.com/spf13/cobra"
	infoCmd "github.com/yourusername/iptool/pkg/commands/info"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "iptool",
		Short: "Small IP/CIDR helper",
	}

	// add subcommands
	root.AddCommand(infoCmd.NewCmdInfo())
	// root.AddCommand(rangeCmd.NewCmdRange()) etc.

	return root
}
