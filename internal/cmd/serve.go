package cmd

import (
	"golang-seed/internal/api"

	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the application server",
		RunE:  run,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func run(cmd *cobra.Command, args []string) error {
	return api.StartServer()
}
