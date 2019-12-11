package service

import "golang-seed/internal/container/model"

type (
	todoService struct {
		TodoModel model.TodoModel
	}

	TodoService interface {
		CreateTodo(todo *model.Todo) error
		GetById(id string) (todo *model.Todo, er error)
	}
)

func NewTodoService(ctx *Context) TodoService {
	return &todoService{
		TodoModel: ctx.TodoModel,
	}
}

func (ts *todoService) CreateTodo(todo *model.Todo) error {
	return ts.TodoModel.CreateTodo(todo)
}

func (ts *todoService) GetById(id string) (todo *model.Todo, er error) {
	return ts.TodoModel.GetById(id)
}
