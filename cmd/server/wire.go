package main

// // go:build wireinject
// //go:build wireinject
// // +build wireinject

// package main

// import (
// 	"io"
// 	"os"

// 	"github.com/Demetrius81/containerized-todo-api/internal/middleware"
// 	"github.com/Demetrius81/containerized-todo-api/internal/repository/postgres"
// 	"github.com/Demetrius81/containerized-todo-api/internal/server"
// 	"github.com/Demetrius81/containerized-todo-api/internal/services/apiservice"
// 	"github.com/Demetrius81/containerized-todo-api/internal/services/logger"
// 	"github.com/google/wire"
// )

// func newWriter() io.Writer {
// 	return os.Stdout
// }

// // Провайдеры для wire
// var (
// 	storageSet = wire.NewSet(
// 		postgres.NewStorage,
// 		wire.Bind(new(apiservice.IStorage), new(*postgres.Storage)),
// 	)

// 	apiSet = wire.NewSet(
// 		apiservice.NewTodoHandlers,
// 		wire.Bind(new(server.ITodoHandlers), new(*apiservice.TodoHandlers)),
// 	)

// 	loggerMiddlewareSet = wire.NewSet(
// 		middleware.NewLoggerMiddleware,
// 		wire.Bind(new(apiservice.ILoggerMiddleware), new(*middleware.LoggerMiddleware)),
// 	)

// 	loggerSet = wire.NewSet(
// 		logger.NewLogger,
// 		wire.Bind(new(middleware.ILogger), new(*logger.Logger)), // интерфейс -> реализация
// 	)
// )

// func InitializeServer(dsn string) (*server.Server, error) {
// 	wire.Build(
// 		postgres.NewDB,
// 		newWriter,
// 		loggerSet,
// 		loggerMiddlewareSet,
// 		storageSet,
// 		apiSet,
// 		server.NewServer,
// 	)
// 	return &server.Server{}, nil
// }
