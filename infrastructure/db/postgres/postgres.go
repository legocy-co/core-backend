package postgres

import (
	"database/sql"
	"fmt"
	"legocy-go/infrastructure/db"
)

type PostrgresConnection struct {
	config *db.ConnectionConfig
}

func (p *PostrgresConnection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.config.Hostname, p.config.Port, p.config.Username, p.config.Password, p.config.DBName)
}

func (p *PostrgresConnection) Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", p.getConnectionString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, err
}
