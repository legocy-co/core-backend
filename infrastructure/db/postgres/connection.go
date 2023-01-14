package postgres

import (
	"fmt"
	config "legocy-go/config"
	d "legocy-go/infrastructure/db"
	entities "legocy-go/infrastructure/db/postgres/entities"

	"github.com/jinzhu/gorm"
)

var postgresConn *PostrgresConnection // private singleton instance

type PostrgresConnection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
}

// Call from outside
func CreateConnection(config *config.DatabaseConfig, db *gorm.DB) (*PostrgresConnection, error) {
	if postgresConn != nil {
		return nil, d.ErrConnectionAlreadyExists
	}
	postgresConn = &PostrgresConnection{config, db}
	return postgresConn, nil
}

func (psql *PostrgresConnection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psql.config.Hostname, psql.config.Port, psql.config.DbUser, psql.config.DbPassword, psql.config.DbName)
}

func (psql *PostrgresConnection) Init() {

	conn, err := gorm.Open("postgres", psql.getConnectionString())
	if err != nil {
		return
	}

	psql.db = conn
	psql.db.LogMode(true)

	psql.db.Debug().AutoMigrate(
		entities.UserPostgres{},
		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},
	)
}

func (psql *PostrgresConnection) GetDB() *gorm.DB {
	return psql.db
}
