package data

import "gorm.io/gorm"

type Storage interface {
	Init() error
	IsReady() bool
	GetDB() *gorm.DB
}
