package api

import (
	"golang-seed/internal/server"

	log "github.com/sirupsen/logrus"
)

func StartServer() (err error) {
	s, err := server.New()
	if err != nil {
		log.Error("failed to create a server", err)
		return
	}

	err = s.ConnectToDabase()
	if err != nil {
		log.Error("failed to create to the database", err)
		return
	}

	err = s.Setup().Start()
	if err != nil {
		log.Error("failed to serve the api", err)
	}
	return
}
