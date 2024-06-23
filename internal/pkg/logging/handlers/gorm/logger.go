package gorm

import (
	"context"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

// Logger logs all queries in plain text
type Logger struct {
	logger *slog.Logger
	config logger.Config
}

func NewLogger(logger *slog.Logger) *Logger {
	return &Logger{
		logger: logger,
		config: getLoggerConfig(),
	}
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.config.LogLevel = level
	return &newLogger
}

func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Info {
		l.logger.Info(msg, slog.Any("data", data))
	}
}

func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Warn {
		l.logger.Warn(msg, slog.Any("data", data))
	}
}

func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Error {
		l.logger.Error(msg, slog.Any("data", data))
	}
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.config.LogLevel <= logger.Silent {
		return
	}

	sql, rows := fc()
	elapsed := time.Since(begin)

	l.logger.Info("trace",
		slog.String("sql", sql),
		slog.Int64("rows_affected", rows),
		slog.Duration("elapsed", elapsed),
	)
}
