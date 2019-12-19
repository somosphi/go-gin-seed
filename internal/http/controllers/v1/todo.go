package v1

import (
	"golang-seed/internal/container"
	"golang-seed/internal/http/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	todoController struct {
		c *container.Container
	}

	TodoController interface {
		controllers.Controller
		createTodo(ctx *gin.Context)
		getAllTodos(ctx *gin.Context)
		getTodo(ctx *gin.Context)
		updateTodo(ctx *gin.Context)
		deleteTodo(ctx *gin.Context)
	}
)

func NewTodoController(c *container.Container) controllers.Controller {
	return &todoController{c}
}

func (tdc *todoController) Register(gr *gin.RouterGroup) {
	todoGroup := gr.Group("/v1/todos")

	todoGroup.POST(
		"",
		tdc.createTodo,
	)
	todoGroup.GET("", tdc.getAllTodos)
	todoGroup.GET("/:todo_id", tdc.getTodo)
	todoGroup.PATCH("/:todo_id", tdc.updateTodo)
	todoGroup.DELETE("/:todo_id", tdc.deleteTodo)
}

func (todoController) createTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"id":   uuid.New(),
		"todo": "created",
	})
}

func (todoController) getTodo(ctx *gin.Context) {}

func (todoController) getAllTodos(ctx *gin.Context) {}

func (todoController) updateTodo(ctx *gin.Context) {}

func (todoController) deleteTodo(ctx *gin.Context) {}
