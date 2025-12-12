package usecase

import (
	"belajar-http/entity"
	"belajar-http/repository"
)

type TodoUsecase interface {
	CreateTodo(title string, description string) (*entity.Todo, error)
	GetAllTodos() ([]entity.Todo, error)
	GetTodoByID(id uint) (*entity.Todo, error)
	UpdateTodo(id uint, title string, description string, done bool) (*entity.Todo, error)
	DeleteTodo(id uint) error
}

type TodoUsecaseImpl struct {
	repo repository.TodoRepository
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo: repo}
}

func (u *todoUsecase) CreateTodo(title, description string) (*entity.Todo, error) {
	todo := &entity.Todo{Title: title, Description: description, Done: false}
	err := u.repo.Create(todo)
	return todo, err
}

func (u *todoUsecase) GetAllTodos() ([]entity.Todo, error) {
	return u.repo.GetAll()
}

func (u *todoUsecase) GetTodoByID(id uint) (*entity.Todo, error) {
	return u.repo.GetByID(id)
}

func (u *todoUsecase) UpdateTodo(id uint, title, description string, completed bool) (*entity.Todo, error) {
	todo, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Description = description
	todo.Done = completed
	err = u.repo.Update(todo)
	return todo, err
}

func (u *todoUsecase) DeleteTodo(id uint) error {
	return u.repo.Delete(id)
}
