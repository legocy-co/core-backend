package db

import "database/sql"

type DBConnection interface {
	Connect() (*sql.DB, error)
}
