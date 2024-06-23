package postgres

import (
	"fmt"
	"github.com/legocy-co/legocy/internal/pkg/config"
	"log/slog"
	"time"

	log "github.com/legocy-co/legocy/internal/pkg/logging/handlers/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresConn *Connection
)

type Connection struct {
	config *config.DatabaseConfig
	db     *gorm.DB
	log    *slog.Logger
}

func New(config *config.DatabaseConfig, log *slog.Logger) (*Connection, error) {
	if postgresConn != nil {
		return nil, ErrConnectionAlreadyExists
	}

	postgresConn = &Connection{config, nil, log}

	if err := postgresConn.Init(); err != nil {
		return nil, err
	}

	return postgresConn, nil
}

func (c *Connection) IsReady() bool {
	return c.db != nil
}

func (c *Connection) Init() error {

	conn, err := gorm.Open(
		postgres.Open(c.getConnectionString()),
		&gorm.Config{
			Logger: log.NewLogger(c.log),
		},
	)

	if err != nil {
		return err
	}

	c.db = conn
	if err := c.applyMigrations(); err != nil {
		return err
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return nil
}

func (c *Connection) GetDB() *gorm.DB {
	return c.db
}

func (c *Connection) GetLogger() *slog.Logger {
	return c.log
}

func (c *Connection) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.config.Hostname, c.config.Port, c.config.DbUser, c.config.DbPassword, c.config.DbName)
}
