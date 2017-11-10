package beego

import (
	"github.com/SinedVonotirah/gopo/persistence/beego/entities"
	"github.com/astaxie/beego/orm"
)

type OrderRepo struct {
	connection orm.Ormer
}

func NewOrderRepo(connection *Connection) *OrderRepo {
	return &OrderRepo{connection.beego}
}

func (repo *OrderRepo) CreateOrder(entity *entities.OrderEntity) (int64, error) {
	repo.connection.QueryM2M(entity, "Products").Add(entity.Products)
	repo.connection.Begin()
	id, err := repo.connection.Insert(entity)
	if err != nil {
		repo.connection.Rollback()
	} else {
		repo.connection.Commit()
	}
	return id, err
}

func (repo *UserRepo) DeleteOrder(id int64) (int64, error) {
	return repo.connection.Delete(&entities.OrderEntity{Id: id})
}

func (repo *UserRepo) UpdateOrder(entity *entities.OrderEntity) (int64, error) {
	return repo.connection.Update(entity)
}

func (repo *OrderRepo) GetOrderById(id int64) (entities.OrderEntity, error) {
	entity := entities.OrderEntity{}
	err := repo.connection.QueryTable("order").Filter("Id", id).RelatedSel().One(&entity)
	repo.connection.LoadRelated(&entity, "Product")
	return entity, err
}

func (repo *OrderRepo) GetOrdersByUserName(userName string) ([]entities.OrderEntity, error) {
	var entities []entities.OrderEntity
	qs := repo.connection.QueryTable("order")
	_, err := qs.Filter("User__Name", userName).All(&entities)
	return entities, err
}

// FIXME
func (repo *OrderRepo) GetOrdersByProductName(productName string) ([]entities.OrderEntity, error) {
	var entities []entities.OrderEntity
	qs := repo.connection.QueryTable("order")
	_, err := qs.Filter("Product__Id", productName).All(&entities)
	return entities, err
}
