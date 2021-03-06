package persistence

import (
	"github.com/SinedVonotirah/gopo/orms/beego/persistence/entities"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(entities.UserGroupEntity),
		new(entities.UserEntity),
		new(entities.OrderEntity),
		new(entities.ProductEntity),
		new(entities.ProductDetailsEntity),
	)
}

type Connection struct {
	Beego orm.Ormer
}

func NewConnection(connectionUrl string) *Connection {
	return &Connection{
		initDb(connectionUrl),
	}
}

func initDb(connectionUrl string) orm.Ormer {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		connectionUrl)

	orm.RunSyncdb("default", false, true)
	//	orm.Debug = true
	return orm.NewOrm()
}
