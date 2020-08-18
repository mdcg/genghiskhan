package cmd

import (
	"fmt"

	"github.com/mdcg/go-port-scanner/scanner"
	"github.com/spf13/cobra"
)

func PrintGSBanner() {
	fmt.Println(
		`
		 ▄████ ▓█████  ███▄    █   ▄████  ██▓  ██████  ██ ▄█▀ ██░ ██  ▄▄▄       ███▄    █ 
		██▒ ▀█▒▓█   ▀  ██ ▀█   █  ██▒ ▀█▒▓██▒▒██    ▒  ██▄█▒ ▓██░ ██▒▒████▄     ██ ▀█   █ 
		▒██░▄▄▄░▒███   ▓██  ▀█ ██▒▒██░▄▄▄░▒██▒░ ▓██▄   ▓███▄░ ▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒
		░▓█  ██▓▒▓█  ▄ ▓██▒  ▐▌██▒░▓█  ██▓░██░  ▒   ██▒▓██ █▄ ░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒
		░▒▓███▀▒░▒████▒▒██░   ▓██░░▒▓███▀▒░██░▒██████▒▒▒██▒ █▄░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░
		░▒   ▒ ░░ ▒░ ░░ ▒░   ▒ ▒  ░▒   ▒ ░▓  ▒ ▒▓▒ ▒ ░▒ ▒▒ ▓▒ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
		 ░   ░  ░ ░  ░░ ░░   ░ ▒░  ░   ░  ▒ ░░ ░▒  ░ ░░ ░▒ ▒░ ▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░
		░ ░   ░    ░      ░   ░ ░ ░ ░   ░  ▒ ░░  ░  ░  ░ ░░ ░  ░  ░░ ░  ░   ▒      ░   ░ ░ 
			 ░    ░  ░         ░       ░  ░        ░  ░  ░    ░  ░  ░      ░  ░         ░ `,
	)
}

func Execute() {
	var host string
	var port int

	rootCmd := &cobra.Command{
		Use: "gengiskhan",
	}

	rootCmd.PersistentFlags().StringVar(
		&host, "host", "localhost",
		"Host you want to scan.",
	)
	rootCmd.PersistentFlags().IntVar(
		&port, "ports", 10000,
		"Number of ports you want to scan.",
	)

	cmdTCPPortScanner := &cobra.Command{
		Use:   "tcp",
		Short: "Initiates scanning of TCP ports.",
		Long: `
			Initiates scanning of TCP ports on a given host.
			Use the --host flags to indicate which host will be scanned,
			and --ports to indicate the number of ports to be scanned.
		`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			scanner.FormatScanReport(scanner.TCPScanner(host, port))
		},
	}

	cmdUDPPortScanner := &cobra.Command{
		Use:   "udp",
		Short: "Initiates scanning of UDP ports.",
		Long: `
			Initiates scanning of UDP ports on a given host.
			Use the --host flags to indicate which host will be scanned,
			and --ports to indicate the number of ports to be scanned.
		`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			scanner.FormatScanReport(scanner.UDPScanner(host, port))
		},
	}

	cmdFullPortScanner := &cobra.Command{
		Use:   "fullScan",
		Short: "Initiates scanning of all ports, whether they are TCP or UDP.",
		Long: `
			Initiates scanning all ports, whether TCP or UDP, on a given host.
			Use the --host flags to indicate which host will be scanned,
			and --ports to indicate the number of ports to be scanned.
		`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			scanner.FormatScanReport(scanner.FullScanner(host, port))
		},
	}

	rootCmd.AddCommand(cmdTCPPortScanner)
	rootCmd.AddCommand(cmdUDPPortScanner)
	rootCmd.AddCommand(cmdFullPortScanner)
	rootCmd.Execute()
}
