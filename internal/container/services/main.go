package services

import (
	"golang-seed/internal/container/models"
)

type Context struct {
	TodoModel models.TodoModel
	PostModel models.PostModel
}

// Creates a new context struct based in the props
func NewContext(props ...interface{}) *Context {
	ctx := &Context{}

	for _, p := range props {
		switch v := p.(type) {
		case models.TodoModel:
			ctx.TodoModel = v.(models.TodoModel)
		case models.PostModel:
			ctx.PostModel = v.(models.PostModel)
		}
	}

	return ctx
}
