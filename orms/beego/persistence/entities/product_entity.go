package entities

import (
	"bytes"
	"fmt"
)

type ProductEntity struct {
	Id             int64
	Name           string
	Orders         []*OrderEntity        `orm:"reverse(many)"`
	ProductDetails *ProductDetailsEntity `orm:"rel(one)"`
}

func (ProductEntity) TableName() string {
	return "product"
}

func (entity ProductEntity) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Product - %+v, %+v", entity.Id, entity.Name))
	for _, order := range entity.Orders {
		buffer.WriteString(order.ToString())
	}
	return buffer.String()
}
