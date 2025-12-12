package repository

import (
	"belajar-http/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *entity.Todo) error
	GetAll() ([]entity.Todo, error)
	GetByID(id uint) (*entity.Todo, error)
	Update(todo *entity.Todo) error
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *entity.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetAll() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *entity.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Todo{}, id).Error
}
