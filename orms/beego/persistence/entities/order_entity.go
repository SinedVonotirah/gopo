package entities

import (
	"bytes"
	"fmt"
)

type OrderEntity struct {
	Id       int64
	Name     string
	User     *UserEntity      `orm:"rel(fk)"`
	Products []*ProductEntity `orm:"rel(m2m)"`
}

func (OrderEntity) TableName() string {
	return "order"
}

func NewOrderEntity(entity *OrderEntity) {
	entity.Id = 0
	entity.Name = "Order"
	entity.User = &UserEntity{Id: 1, Name: "name", UserGroup: &UserGroupEntity{Id: 1, Name: "Group1"}}
	entity.Products = []*ProductEntity{
		&ProductEntity{Id: 1, Name: "product", ProductDetails: &ProductDetailsEntity{Id: 1, Description: "desc"}},
	}
}

func (entity OrderEntity) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Order - %+v, %+v, %+v", entity.Id, entity.Name))
	buffer.WriteString(entity.User.ToString())
	/*for _, product := range entity.Products {
		buffer.WriteString(product.ToString())
	}*/
	return buffer.String()
}
