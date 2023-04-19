package database

import (
	"gorm.io/gorm"
)

type Database interface {
	connect(dns string) *gorm.DB
}
