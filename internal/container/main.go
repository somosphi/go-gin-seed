package container

import (
	"golang-seed/internal/container/models"
	"golang-seed/internal/container/services"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		TodoService services.TodoService
		PostService services.PostService
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(cfg *ContainerConfig) *Container {
	svcCtx := services.NewContext(
		models.NewTodoMysql(cfg.Database),
		models.NewPostModel(cfg.Database),
	)

	return &Container{
		TodoService: services.NewTodoService(svcCtx),
		PostService: services.NewPostService(svcCtx),
	}
}
