package app

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/internal/pkg/logging"
	"log/slog"
)

var (
	logger *slog.Logger
)

func (a *App) GetLogger() *slog.Logger {

	if logger == nil {
		cfg := config.GetAppConfig()
		logger = logging.SetupLogger(cfg.Environment)
	}

	return logger
}
