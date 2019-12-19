package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	Todo struct {
		gorm.Model
		Title     string `gorm:"default:'Empty Title'" json:"title"`
		Completed int    `gorm:"type:tinyint(1);default:0" json:"completed"`
	}

	todoMysql struct {
		db *gorm.DB
	}

	TodoModel interface {
		CreateTodo(t *Todo) error
		GetById(id string) (todo *Todo, er error)
	}
)

// Create new Todo table
func NewTodoMysql(db *gorm.DB) (tm TodoModel) {
	db.AutoMigrate(&Todo{})
	return &todoMysql{db}
}

// BeforeCreate add uuid as Todo.ID
func (*Todo) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New())
	return err
}

// Insert a new todo in the database
func (tm *todoMysql) CreateTodo(todo *Todo) error {
	tm.db.Create(todo)
	return nil
}

// Find a todo by ID
func (tm *todoMysql) GetById(id string) (todo *Todo, er error) {
	todo = &Todo{}
	tm.db.First(todo, "id = ?", id)
	return
}
