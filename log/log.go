package log

import (
	"log/slog"
	"os"

	"framework/config"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(config *config.Config) *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(config.GetInt("log.level"))}))
	return &Logger{logger}
}

func (l *Logger) WithModule(module string) *Logger {
	return &Logger{l.Logger.With(slog.String("module", module))}
}

func Any(key string, value any) slog.Attr {
	return slog.Any(key, value)
}

func String(key string, value string) slog.Attr {
	return slog.String(key, value)
}

func Int(key string, value int) slog.Attr {
	return slog.Int(key, value)
}
