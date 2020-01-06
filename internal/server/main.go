package server

import (
	"fmt"
	"golang-seed/internal/container"
	"golang-seed/internal/helpers"
	"golang-seed/internal/http"
	log "golang-seed/internal/http/middlewares/log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Config    *Config
	Container *container.Container
	engine    *gin.Engine
	db        *gorm.DB
}

func createDatabase() (db *gorm.DB, err error) {
	dbConfig, err := helpers.NewDbConfig()
	if err == nil && dbConfig != nil {
		db, err = helpers.GetDatabase(dbConfig)
	}

	return
}

func New() (s *Server, err error) {
	s = &Server{}
	s.Config, err = InitConfig()
	return
}

func (s *Server) ConnectToDabase() error {
	db, err := createDatabase()
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Server) Setup() *Server {
	s.Container = container.NewContainer(
		container.NewContainerConfig(s.db),
	)

	s.engine = gin.New()

	log.AttachReqResLogger(s.engine, log.New())
	s.engine.Use(gin.Recovery())

	http.SetupRoutes(s.engine, s.Container)
	return s
}

func (s *Server) Start() error {
	log.Logger.Infof("server listening on localhost:%d\n", s.Config.Port)
	return s.engine.Run(fmt.Sprintf(":%d", s.Config.Port))
}

func (s *Server) CloseDB() {
	s.db.Close()
}
