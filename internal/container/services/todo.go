package services

import "golang-seed/internal/container/models"

type (
	todoService struct {
		TodoModel models.TodoModel
	}

	TodoService interface {
		CreateTodo(todo *models.Todo) error
		GetById(id string) (todo *models.Todo, er error)
	}
)

func NewTodoService(ctx *Context) TodoService {
	return &todoService{
		TodoModel: ctx.TodoModel,
	}
}

func (ts *todoService) CreateTodo(todo *models.Todo) error {
	return ts.TodoModel.CreateTodo(todo)
}

func (ts *todoService) GetById(id string) (todo *models.Todo, er error) {
	return ts.TodoModel.GetById(id)
}
