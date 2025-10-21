package middleware

import (
	"net/http"
	"time"

	"github.com/Demetrius81/containerized-todo-api/internal/services/logger"
)

type ILogger interface {
	Log(w logger.ILoggingResponseWriter, r *http.Request, elapsedTime time.Duration)
}

type LoggerMiddleware struct {
	logger ILogger
}

func NewLoggerMiddleware(logger ILogger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger: logger,
	}
}

// LoggingMiddleware логирует метод, путь, статус и время обработки.
func (cl *LoggerMiddleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		lrw := &WrapperResponseWriter{ResponseWriter: writer}

		next.ServeHTTP(lrw, req)

		elapsedTime := time.Since(startTime)
		cl.logger.Log(lrw, req, elapsedTime)

	})
}

// ResponseWriter wrapper for catch status code
type WrapperResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *WrapperResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *WrapperResponseWriter) GetStatusCode() int {
	return rw.statusCode
}
