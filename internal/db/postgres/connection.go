package postgres

import (
	"fmt"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	"log"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateConnection(config *config.DatabaseConfig, db *gorm.DB) (d.DataBaseConnection, error) {
	if postgresConn != nil {
		return nil, d.ErrConnectionAlreadyExists
	}
	postgresConn = &PostgresConnection{config, db}
	return postgresConn, nil
}

var postgresConn *PostgresConnection // private singleton instance
type PostgresConnection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
}

func (psql *PostgresConnection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psql.config.Hostname, psql.config.Port, psql.config.DbUser, psql.config.DbPassword, psql.config.DbName)
}

func (psql *PostgresConnection) Init() {
	dsn := psql.getConnectionString()
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("Error connecting to database! %v", err.Error())
		panic(err)
	}

	psql.db = conn

	err = psql.db.Debug().AutoMigrate(
		entities.UserPostgres{},
		entities.UserPostgresImage{},

		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},

		entities.CurrencyPostgres{},
		entities.LocationPostgres{},
		entities.MarketItemPostgres{},
	)

	if err != nil {
		log.Fatalln(fmt.Sprintf("[Postgres] %v", err.Error()))
	}
}

func (psql *PostgresConnection) GetDB() *gorm.DB {
	return psql.db
}
