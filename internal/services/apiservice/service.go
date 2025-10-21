package apiservice

import (
	"net/http"

	"github.com/Demetrius81/containerized-todo-api/internal/domain"
	"github.com/gorilla/mux"
)

type TodoHandlers struct {
	storage IStorage
	logger  ILoggerMiddleware
}

type ILoggerMiddleware interface {
	LogMiddleware(next http.Handler) http.Handler
}

type IStorage interface {
	Create(todo domain.Todo) (domain.Todo, error)
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (domain.Todo, error)
	Update(id uint, todo domain.Todo) (domain.Todo, error)
	Delete(id uint) error
}

func NewTodoHandlers(storage IStorage, logger ILoggerMiddleware) *TodoHandlers {
	return &TodoHandlers{
		storage: storage,
		// mux:     mux.NewRouter(),
		logger: logger,
	}
}

func (s *TodoHandlers) RegisterHandlers(mux *mux.Router) {
	mux.HandleFunc("/todos", s.HandlerGetTodos).Methods("GET")
	mux.HandleFunc("/todos/{id:[0-9]+}", s.HandlerGetTodo).Methods("GET")
	mux.HandleFunc("/todos", s.HandlerCreateTodo).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", s.HandlerUpdateTodo).Methods("PUT")
	mux.HandleFunc("/todos/{id:[0-9]+}", s.HandlerDeleteTodo).Methods("DELETE")
	mux.Use(s.logger.LogMiddleware)
}
