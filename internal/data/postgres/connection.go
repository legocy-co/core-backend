package postgres

import (
	"fmt"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"github.com/legocy-co/legocy/pkg/logging"
	log "github.com/sirupsen/logrus"
	"time"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(config *config.DatabaseConfig, db *gorm.DB) (d.DBConn, error) {
	if postgresConn != nil {
		return nil, d.ErrConnectionAlreadyExists
	}
	postgresConn = &Connection{config, db}
	return postgresConn, nil
}

var postgresConn *Connection // private singleton instance
type Connection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
}

func (psql *Connection) IsReady() bool {
	return psql.db != nil
}

func (psql *Connection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psql.config.Hostname, psql.config.Port, psql.config.DbUser, psql.config.DbPassword, psql.config.DbName)
}

func (psql *Connection) Init() {
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
		entities.MarketItemLikePostgres{},

		entities.UserImagePostgres{},

		entities.UserReviewPostgres{},
		entities.UserLegoSetPostgres{},

		entities.LegoSetImagePostgres{},
	)

	if err != nil {
		log.Fatalln(fmt.Sprintf("[Postgres] %v", err.Error()))
	}

	sqlDB, err := psql.db.DB()
	if err != nil {
		log.Fatalln(fmt.Sprintf("[Postgres] %v", err.Error()))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
}

func (psql *Connection) GetDB() *gorm.DB {
	return psql.db
}
