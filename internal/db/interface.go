package db

import "gorm.io/gorm"

type DataBaseConnection interface {
	Init()
	GetDB() *gorm.DB
}
