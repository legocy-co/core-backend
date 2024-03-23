package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/legocy-co/legocy/config"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/fixtures"
	"github.com/legocy-co/legocy/pkg/kafka"
	log "github.com/sirupsen/logrus"
	"time"
)

func New() *App {

	godotenv.Load()

	app := App{}

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
	if cfg == nil {
		log.Fatalln("empty data config")
	}
	app.setDatabase(dbCfg)

	// Fixtures
	go func(load bool) {
		if !load {
			return
		}
		fixtures.LoadLegoSeries(app.GetLegoSeriesRepo())
		fixtures.LoadLegoSets(app.GetLegoSetRepo(), app.GetLegoSeriesRepo())
	}(dbCfg.LoadFixtures)

	// Check all deps
	if !app.isReady() {
		panic("Some dependencies failed to inject")
	}

	return &app
}

type App struct {
	database d.DataBaseConnection
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
