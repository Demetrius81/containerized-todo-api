package logger

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{logger: slog.New(slog.NewTextHandler(w, nil))}
}

// Log implements ILogger.
func (l *Logger) Log(w ILoggingResponseWriter, r *http.Request, elapsedTime time.Duration) {
	l.logger.Info("Request handled > ",
		slog.String("method", r.Method),
		slog.String("host", r.Host),
		slog.String("path", r.URL.Path),
		slog.String("remote_addr", r.RemoteAddr),
		slog.Int("status", w.GetStatusCode()),
		slog.Duration("latency", elapsedTime),
	)
}

type ILoggingResponseWriter interface {
	GetStatusCode() int
}
