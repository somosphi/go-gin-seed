package services

import (
	"golang-seed/internal/container/models"
)

type (
	postService struct {
		PostModel models.PostModel
	}

	PostService interface {
		CreatePost(post *models.Post) error
	}
)

func NewPostService(ctx *Context) PostService {
	return &postService{
		PostModel: ctx.PostModel,
	}
}

func (ps *postService) CreatePost(post *models.Post) error {
	return nil
}
