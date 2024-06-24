package logging

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	"log/slog"
	"os"
)

func SetupLogger(env config.Environment) *slog.Logger {

	var log *slog.Logger

	switch env {
	case config.EnvDevelopment:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProduction:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
