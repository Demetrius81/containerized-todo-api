package postgres

import (
	"errors"
	"fmt"

	"github.com/Demetrius81/containerized-todo-api/internal/domain"
	"github.com/Demetrius81/containerized-todo-api/internal/repository/postgres/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	fmt.Println(">>>")
	fmt.Println(dialector)
	fmt.Println(">>>")

	return gorm.Open(dialector, &gorm.Config{})
}

type Storage struct {
	db *gorm.DB
}

// NewStorage создает новое хранилище и выполняет миграции.
func NewStorage(db *gorm.DB) (*Storage, error) {
	if err := migrations.RunMigrations(db); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

// CRUD методы
func (s *Storage) Create(todo domain.Todo) (domain.Todo, error) {
	result := s.db.Create(&todo)

	if err := result.Error; err != nil {
		return domain.Todo{}, err
	}

	if result.RowsAffected == 0 {
		err := errors.New("can not create todo")
		return domain.Todo{}, err
	}

	return todo, nil
}

func (s *Storage) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	result := s.db.Find(&todos)
	if err := result.Error; err != nil {
		return []domain.Todo{}, err
	}

	return todos, nil
}

func (s *Storage) GetByID(id uint) (domain.Todo, error) {
	var todo domain.Todo
	result := s.db.First(&todo, id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Todo{}, err
		}
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *Storage) Update(id uint, todo domain.Todo) (domain.Todo, error) {
	var existingTodo domain.Todo
	result := s.db.First(&existingTodo, id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Todo{}, err
		}
		return domain.Todo{}, err
	}

	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	existingTodo.Done = todo.Done
	result = s.db.Save(&existingTodo)
	if err := result.Error; err != nil {
		return domain.Todo{}, err
	}

	return existingTodo, nil
}

func (s *Storage) Delete(id uint) error {
	result := s.db.Delete(&domain.Todo{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
