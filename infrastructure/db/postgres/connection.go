package postgres

import (
	"fmt"
	config "legocy-go/config"
	d "legocy-go/infrastructure/db"
	entities "legocy-go/infrastructure/db/postgres/entities"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var postgresConn *PostgresConnection // private singleton instance
type PostgresConnection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
}

func CreateConnection(config *config.DatabaseConfig, db *gorm.DB) (*PostgresConnection, error) {
	if postgresConn != nil {
		return nil, d.ErrConnectionAlreadyExists
	}
	postgresConn = &PostgresConnection{config, db}
	return postgresConn, nil
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

	psql.db.Debug().AutoMigrate(
		entities.UserPostgres{},
		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},
		entities.CurrencyPostgres{},
		entities.LocationPostgres{},
		entities.MarketItemPostgres{},
	)
}

func (psql *PostgresConnection) GetDB() *gorm.DB {
	return psql.db
}
