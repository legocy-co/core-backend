package postgres

import (
	"fmt"
	config "legocy-go/config"
	entities "legocy-go/infrastructure/db/postgres/entities"

	"github.com/jinzhu/gorm"
)

type PostrgresConnection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
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
		entities.LegoSeriesPostgres{},
		entities.LegoSetPostgres{},
	)
}

func (psql *PostrgresConnection) GetDB() *gorm.DB {
	defer psql.db.Close()
	return psql.db
}
