package cmd

import "github.com/spf13/cobra"

func PortScannerRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "Port Scanner",
		Short: "List Ports",
		Long:  "List all TCP Opened ports",
	}
}
