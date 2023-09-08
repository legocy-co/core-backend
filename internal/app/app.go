package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	"legocy-go/internal/fixtures"
	"legocy-go/pkg/kafka"
	"log"
	"time"
)

type App struct {
	database d.DataBaseConnection
}

func (a *App) isReady() bool {
	dbReady := a.database.IsReady()

	if !dbReady {
		logrus.Error("DB Connection Failed...")
		return false
	}

	logrus.Info("Checking Kafka...")

	ctx, cf := context.WithTimeout(context.Background(), time.Second*3)
	defer cf()

	kafkaReady := kafka.IsKafkaConnected(ctx)
	if !kafkaReady {
		logrus.Error("Kafka Connection Failed...")
		return false
	}

	return true
}

func New(configFilepath string) *App {

	app := App{}

	// Load config
	err := config.SetupFromJSON(configFilepath)
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
		log.Fatalln("empty db config")
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
