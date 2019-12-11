package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	Post struct {
		gorm.Model
		Title      string `gorm:"not null"`
		Content    string `gorm:"not null"`
		AuthorName string `gorm:"default:\"Desconhecido\""`
	}

	postMysql struct {
		db *gorm.DB
	}

	PostModel interface {
		CreatePost(post *Post) error
	}
)

func NewPostModel(db *gorm.DB) PostModel {
	db.AutoMigrate(&Post{})
	return &postMysql{db}
}

// BeforeCreate add uuid as Post.ID
func (t *Post) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New())
	return err
}

func (pm *postMysql) CreatePost(post *Post) error {
	pm.db.Create(post)
	return nil
}
