// pkg/commands/info/info.go
package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"inet.af/netaddr"

	"github.com/kavishbaghel/iptool/internal/ip"
)

func NewCmdInfo() *cobra.Command {
	var outFmt string

	cmd := &cobra.Command{
		Use:   "info <cidr>",
		Short: "Show network info for a CIDR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := netaddr.ParseIPPrefix(args[0])
			if err != nil {
				return fmt.Errorf("invalid CIDR: %w", err)
			}

			res, err := ip.Info(p)
			if err != nil {
				return err
			}

			switch outFmt {
			case "json":
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				return enc.Encode(res)
			case "text", "":
				printText(res)
				return nil
			default:
				return fmt.Errorf("unknown format %q", outFmt)
			}
		},
	}

	cmd.Flags().StringVarP(&outFmt, "format", "f", "text", "output format: text|json")
	return cmd
}

func printText(r ip.InfoResult) {
	fmt.Printf("Network:   %s\n", r.Network)
	fmt.Printf("Prefix:    /%d (bits: %d)\n", r.PrefixLen, r.Bits)
	fmt.Printf("Range:     %s - %s\n", r.From, r.To)
	if r.Total > 0 {
		fmt.Printf("Total:     %d\n", r.Total)
	}
}
