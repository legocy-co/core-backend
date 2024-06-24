package gorm

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	"gorm.io/gorm/logger"
)

func getLoggerConfig() logger.Config {
	var cfg logger.Config

	appCfg := config.GetAppConfig()
	switch appCfg.Environment {
	case config.EnvDevelopment:
		cfg = logger.Config{
			LogLevel: logger.Info,
		}
	case config.EnvProduction:
		cfg = logger.Config{
			LogLevel: logger.Error,
		}
	default:
		cfg = logger.Config{
			LogLevel: logger.Silent,
		}
	}

	return cfg
}
