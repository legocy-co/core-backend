package db

import "github.com/jinzhu/gorm"

type DataBaseConnection interface {
	Init()
	GetDB() *gorm.DB
}
