package http

import (
	"golang-seed/internal/container"
	"golang-seed/internal/http/controllers"
	v1 "golang-seed/internal/http/controllers/v1"

	"github.com/gin-gonic/gin"
)

func _loadControllers(c *container.Container) []controllers.Controller {
	return []controllers.Controller{
		v1.NewTodoController(c),
		v1.NewPostController(c),
	}
}

func SetupRoutes(e *gin.Engine, c *container.Container) {
	api := e.Group("/")

	controls := _loadControllers(c)

	for _, ctr := range controls {
		ctr.Register(api)
	}

	api.GET("/_hc", controllers.HealthCheck)
}
