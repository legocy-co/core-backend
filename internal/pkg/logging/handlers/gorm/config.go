package gorm

import (
	"github.com/legocy-co/legocy/internal/pkg/config"
	"gorm.io/gorm/logger"
)

func getLoggerConfig() logger.Config {
	appCfg := config.GetAppConfig()
	switch appCfg.Environment {
	case config.EnvDevelopment:
		return logger.Config{
			LogLevel: logger.Info,
		}
	case config.EnvProduction:
		return logger.Config{
			LogLevel: logger.Error,
		}
	default:
		return logger.Config{
			LogLevel: logger.Silent,
		}
	}
}
