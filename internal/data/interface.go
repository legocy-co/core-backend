package data

import "gorm.io/gorm"

type DBConn interface {
	Init()
	IsReady() bool
	GetDB() *gorm.DB
}
