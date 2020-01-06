package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	Todo struct {
		ID        uuid.UUID `gorm:"primary_key;type:varchar(64);" json:"id"`
		Title     string    `gorm:"default:'Empty Title'" json:"title"`
		Completed int       `gorm:"type:tinyint(1);default:0" json:"completed"`
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

func (t *Todo) String() string {
	return fmt.Sprintf("Title: %s\nCompleted: %b", t.Title, t.Completed)
}

func (t *Todo) ToJsonMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        t.ID,
		"title":     t.Title,
		"completed": t.Completed,
	}
}

// BeforeCreate add uuid as Todo.ID
func (t *Todo) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	err := scope.SetColumn("ID", id)
	if err == nil {
		t.ID = id
	}
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
