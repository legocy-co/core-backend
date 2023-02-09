package db

import "gorm.io/gorm"

type DataBaseConnection interface {
	Init()
	IsReady() bool
	GetDB() *gorm.DB
}
