package cmd

import (
	log "golang-seed/internal/http/middlewares/log"
	"golang-seed/internal/server"

	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the application server",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			s, err := server.New()
			if err != nil {
				log.Logger.Error("failed to create a server", err)
				return
			}

			err = s.ConnectToDabase()
			if err != nil {
				log.Logger.Error("failed to create to the database", err)
				return
			}
			defer s.CloseDB()

			err = s.Setup().Start()
			if err != nil {
				log.Logger.Error("failed to serve the api", err)
			}
			return
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
	log.New()
}
