package v1

import (
	"encoding/json"
	"golang-seed/internal/container"
	"golang-seed/internal/container/model"
	"golang-seed/internal/http/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
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

	todoGroup.POST("", tdc.createTodo)
	todoGroup.GET("", tdc.getAllTodos)
	todoGroup.GET("/:todo_id", tdc.getTodo)
	todoGroup.PATCH("/:todo_id", tdc.updateTodo)
	todoGroup.DELETE("/:todo_id", tdc.deleteTodo)
}

func (tc *todoController) createTodo(ctx *gin.Context) {
	var body *model.Todo
	rawData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse data",
			"err":     err,
		})
		return
	}

	_ = json.Unmarshal(rawData, &body)
	if body == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse data",
			"err":     "invlid request body",
		})
		return
	}

	err = tc.c.TodoService.CreateTodo(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create todo",
			"err":     err,
		})
		return
	}

	resJson := body.ToJsonMap()

	ctx.JSON(http.StatusCreated, resJson)
}

func (todoController) getTodo(ctx *gin.Context) {}

func (todoController) getAllTodos(ctx *gin.Context) {}

func (todoController) updateTodo(ctx *gin.Context) {}

func (todoController) deleteTodo(ctx *gin.Context) {}
