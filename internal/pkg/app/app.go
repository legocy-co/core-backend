package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/internal/pkg/kafka"
	log "github.com/sirupsen/logrus"
	"time"
)

func New() (*App, error) {

	godotenv.Load()

	app := &App{}

	// Load config json
	log.SetFormatter(&log.JSONFormatter{})

	// Load config
	err := config.SetupFromEnv()
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Config] %v", err.Error()))
	}

	cfg := config.GetAppConfig()
	if cfg == nil {
		log.Fatalln("Error getting app config")
	}

	//Database
	dbCfg := config.GetDBConfig()
	if err := app.setDatabase(dbCfg); err != nil {
		return nil, err
	}

	// Check all deps
	if !app.isReady() {
		panic("Some dependencies failed to inject")
	}

	return app, nil
}

type App struct {
	database d.Storage
}

func (a *App) isReady() bool {
	dbReady := a.database.IsReady()

	if !dbReady {
		log.Error("DB Connection Failed...")
		return false
	}

	log.Info("Checking Kafka...")

	ctx, cf := context.WithTimeout(context.Background(), time.Second*3)
	defer cf()

	kafkaReady := kafka.IsKafkaConnected(ctx)
	if !kafkaReady {
		log.Error("Kafka Connection Failed...")
		return false
	}

	return true
}
