package v1

import (
	"golang-seed/internal/container"
	"golang-seed/internal/http/controllers"

	"github.com/gin-gonic/gin"
)

type (
	postController struct {
		*container.Container
	}

	PostController interface {
		controllers.Controller
		createPost(ctx *gin.Context)
	}
)

func NewPostController(c *container.Container) controllers.Controller {
	return &postController{}
}

func (pc *postController) Register(g *gin.RouterGroup) {
	group := g.Group("/v1/post")

	group.POST("", pc.createPost)
}

func (pc *postController) createPost(ctx *gin.Context) {}
