package inmemory

import (
	"time"

	"github.com/Demetrius81/containerized-todo-api/internal/domain"
	"gorm.io/gorm"
)

type InMemoryStorage struct {
	Db     map[uint]domain.Todo
	NextID uint
}

// NewStorage создает новое хранилище и выполняет миграции.
func NewInMemoryStorage() (*InMemoryStorage, error) {
	return &InMemoryStorage{
		Db:     make(map[uint]domain.Todo),
		NextID: 1,
	}, nil
}

// CRUD методы
func (s *InMemoryStorage) Create(todo domain.Todo) (domain.Todo, error) {
	if todo.Title == "" || todo.Description == "" {
		return domain.Todo{}, gorm.ErrInvalidData
	}

	todo.ID = s.NextID
	s.NextID++
	todo.Done = false
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	s.Db[todo.ID] = todo
	return todo, nil
}

func (s *InMemoryStorage) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	for _, todo := range s.Db {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *InMemoryStorage) GetByID(id uint) (domain.Todo, error) {
	result, ok := s.Db[id]
	if !ok {
		return domain.Todo{}, gorm.ErrRecordNotFound
	}
	return result, nil
}

func (s *InMemoryStorage) Update(id uint, todo domain.Todo) (domain.Todo, error) {
	if todo.Title == "" || todo.Description == "" {
		return domain.Todo{}, gorm.ErrInvalidData
	}
	existingTodo, ok := s.Db[id]
	if !ok {
		return domain.Todo{}, gorm.ErrRecordNotFound
	}

	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	existingTodo.Done = todo.Done

	s.Db[existingTodo.ID] = existingTodo

	return existingTodo, nil
}

func (s *InMemoryStorage) Delete(id uint) error {
	delete(s.Db, id)
	return nil
}
