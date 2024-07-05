package repository

import (
	todoapp "todo-app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GetUser(username, password string) (todoapp.User,error)
}
type TodoList interface {
	Create(userId int, list todoapp.TodoList) (int, error)
}
type TodoItem interface {}


type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
	}
}

