package postgres

import (
	"fmt"
	"legocy-go/internal/config"
	d "legocy-go/internal/db"
	postgres2 "legocy-go/internal/db/postgres/entities"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateConnection(config *config.DatabaseConfig, db *gorm.DB) (*PostgresConnection, error) {
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
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to database!", err)
		return
	}

	psql.db = conn

	err = psql.db.Debug().AutoMigrate(
		postgres2.UserPostgres{},
		postgres2.LegoSeriesPostgres{},
		postgres2.LegoSetPostgres{},
		postgres2.CurrencyPostgres{},
		postgres2.LocationPostgres{},
		postgres2.MarketItemPostgres{},
	)

	if err != nil {
		panic(err)
	}
}

func (psql *PostgresConnection) GetDB() *gorm.DB {
	return psql.db
}
