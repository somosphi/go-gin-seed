package logger

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
)

var Logger *logrus.Logger
var once sync.Once

func getLogLevel() logrus.Level {
	switch viper.GetString("api.logLevel") {
	case "INFO":
		return logrus.InfoLevel
	case "FATAL":
		return logrus.FatalLevel
	}
	// default return is DEBUG level
	return logrus.DebugLevel
}

func New() *logrus.Logger {
	once.Do(func() {
		Logger = logrus.New()

		Logger.SetFormatter(&logrus.JSONFormatter{})
		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(getLogLevel())
	})

	return Logger
}

func AttachReqResLogger(e *gin.Engine, l *logrus.Logger) {
	e.Use(ginlogrus.Logger(l))
}
