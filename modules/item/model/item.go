package model

import (
	"errors"
	"todo/common"
)

type ToDoItem struct {
	common.SQLModel
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      *ItemStatus `json:"status"`
}

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("this item was deleted")
)

func (ToDoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	ID          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (TodoItemCreation) TableName() string {
	return ToDoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       string  `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      string  `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return ToDoItem{}.TableName()
}
