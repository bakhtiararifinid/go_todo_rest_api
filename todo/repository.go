package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Todo, error)
	GetById(id int) (Todo, error)
	AddTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodo(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func CreateNewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return todos, err
	} else {
		return todos, nil
	}
}

func (r *repository) GetById(id int) (Todo, error) {
	var todo Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return todo, err
	} else {
		return todo, nil
	}
}

func (r *repository) AddTodo(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	} else {
		return todo, nil
	}
}

func (r *repository) UpdateTodo(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	} else {
		return todo, nil
	}
}

func (r *repository) DeleteTodo(todo Todo) (Todo, error) {
	err := r.db.Delete(&todo).Error
	if err != nil {
		return todo, err
	} else {
		return todo, nil
	}
}
