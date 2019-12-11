package service

import "golang-seed/internal/container/model"

type Context struct {
	TodoModel model.TodoModel
	PostModel model.PostModel
}

// Creates a new context struct based in the props
func NewContext(props ...interface{}) *Context {
	ctx := &Context{}

	for _, p := range props {
		switch v := p.(type) {
		case model.TodoModel:
			ctx.TodoModel = v.(model.TodoModel)
		case model.PostModel:
			ctx.PostModel = v.(model.PostModel)
		}
	}

	return ctx
}
