package logger

import (
	"sync"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

// GetLogger returns the singleton instance of the logger
func GetLogger() (*zap.Logger, error) {
	var err error
	once.Do(func() {
		logger, err = zap.NewProduction()
	})
	return logger, err
}

func AttachReqResLogger(e *gin.Engine, log *zap.Logger) {
	if log == nil {
		log, _ = GetLogger()
	}

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	e.Use(ginzap.Ginzap(log, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	e.Use(ginzap.RecoveryWithZap(log, true))
}
