package xorm

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Connection struct {
	xorm *xorm.Engine
}

func NewConnection(connectionUrl string) *Connection {
	return &Connection{
		initDb(connectionUrl),
	}
}

func initDb(connectionUrl string) *xorm.Engine {

	db, err := xorm.NewEngine("postgres", connectionUrl)

	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("xorm.Open error")
	}
	return db
}
