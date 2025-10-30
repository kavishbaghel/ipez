/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rangeCmd represents the range command
var rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Return the starting ip address, ending ip address, and the total number of hosts within the CIDR.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip.getCidrRange(args[1])
	},
}

func init() {
	rootCmd.AddCommand(rangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
