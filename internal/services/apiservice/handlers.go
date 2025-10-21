package apiservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Demetrius81/containerized-todo-api/internal/domain"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (s *TodoHandlers) HandlerGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.storage.GetAll()
	if err != nil {
		http.Error(w, "can not get data from server", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *TodoHandlers) HandlerGetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetVarByKey(r, "id")
	if err != nil {
		http.Error(w, "id parse error", http.StatusBadRequest)
		return
	}
	todo, err := s.storage.GetByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, "can not get data from server", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *TodoHandlers) HandlerCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "body decoding error", http.StatusBadRequest)
		return
	}
	_, err = s.storage.Create(todo)
	if err != nil {
		// TODO gorm dependency
		if err == gorm.ErrInvalidData {
			http.Error(w, "todo not found", http.StatusBadRequest)
			return
		}
		http.Error(w, "can not save data to server", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *TodoHandlers) HandlerUpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetVarByKey(r, "id")
	if err != nil {
		http.Error(w, "id parse error", http.StatusBadRequest)
		return
	}
	var todo domain.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "body decoding error", http.StatusBadRequest)
		return
	}
	rez, err := s.storage.Update(id, todo)
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "todo not found", http.StatusNotFound)
			return
		}
		http.Error(w, "can not update data on server", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(rez)
	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *TodoHandlers) HandlerDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := GetVarByKey(r, "id")
	if err != nil {
		http.Error(w, "id parse error", http.StatusBadRequest)
		return
	}
	err = s.storage.Delete(id)
	if err != nil {
		http.Error(w, "can not delete data from server", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetVarByKey(r *http.Request, key string) (uint, error) {
	queryValues := mux.Vars(r)
	idStr, ok := queryValues["id"]
	if !ok {
		return 0, fmt.Errorf("missing 'id' query parameter")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("wrong 'id' query parameter")
	}

	return uint(id), nil
}
