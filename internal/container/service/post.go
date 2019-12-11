package service

import "golang-seed/internal/container/model"

type (
	postService struct {
		PostModel model.PostModel
	}

	PostService interface {
		CreatePost(post *model.Post) error
	}
)

func NewPostService(ctx *Context) PostService {
	return &postService{
		PostModel: ctx.PostModel,
	}
}

func (ps *postService) CreatePost(post *model.Post) error {
	return nil
}
