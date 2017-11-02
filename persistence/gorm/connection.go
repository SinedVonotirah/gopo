package gorm

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Connection struct {
	gorm *gorm.DB
}

func NewConnection(connectionUrl string) *Connection {
	return &Connection{
		initDb(connectionUrl),
	}
}

func initDb(connectionUrl string) *gorm.DB {
	db, err := gorm.Open("postgres", connectionUrl)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
	}
	db.LogMode(true)
	return db
}
