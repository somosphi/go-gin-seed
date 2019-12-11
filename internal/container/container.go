package container

import (
	m "golang-seed/internal/container/model"
	s "golang-seed/internal/container/service"

	"github.com/jinzhu/gorm"
)

type (
	ContainerConfig struct {
		Database *gorm.DB
	}

	Container struct {
		TodoService s.TodoService
		PostService s.PostService
	}
)

func NewContainerConfig(db *gorm.DB) *ContainerConfig {
	return &ContainerConfig{db}
}

func NewContainer(cfg *ContainerConfig) *Container {
	svcCtx := s.NewContext(
		m.NewTodoMysql(cfg.Database),
		m.NewPostModel(cfg.Database),
	)

	return &Container{
		TodoService: s.NewTodoService(svcCtx),
		PostService: s.NewPostService(svcCtx),
	}
}
