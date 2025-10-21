package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type ITodoHandlers interface {
	RegisterHandlers(mux *mux.Router)
}

type Server struct {
	Mux *mux.Router
}

// New создаёт новый Server-сервер.
func NewServer(handler ITodoHandlers) *Server {
	mux := mux.NewRouter()
	handler.RegisterHandlers(mux)
	return &Server{
		Mux: mux,
	}
}

// Start запускает HTTP-сервер.
func (s *Server) Start(addr string) error {
	slog.Info(fmt.Sprintf("Server starting on %s", addr))
	return http.ListenAndServe(addr, s.Mux)
}
