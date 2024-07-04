package service

import (
	todoapp "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(username, password string)  (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {}
type TodoItem interface {}


type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

