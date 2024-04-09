package postgres

import (
	"fmt"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/logging"
	log "github.com/sirupsen/logrus"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (p *PostgresConnection) IsReady() bool {
	return p.db != nil
}

func (psql *PostgresConnection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psql.config.Hostname, psql.config.Port, psql.config.DbUser, psql.config.DbPassword, psql.config.DbName)
}

func (psql *PostgresConnection) Init() {
	dsn := psql.getConnectionString()
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logging.NewGORMLogger(),
	})
	if err != nil {
		fmt.Printf("Error connecting to database! %v", err.Error())
		panic(err)
	}

	psql.db = conn

	err = psql.db.Debug().AutoMigrate(
		entities.UserPostgres{},
		entities.UserImagePostgres{},

		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},
		entities.LegoSetValuationPostgres{},

		entities.MarketItemPostgres{},
		entities.MarketItemImagePostgres{},

		entities.UserImagePostgres{},

		entities.UserReviewPostgres{},
		entities.UserLegoSetPostgres{},

		entities.LegoSetImagePostgres{},
	)

	if err != nil {
		log.Fatalln(fmt.Sprintf("[Postgres] %v", err.Error()))
	}
}

func (psql *PostgresConnection) GetDB() *gorm.DB {
	return psql.db
}
