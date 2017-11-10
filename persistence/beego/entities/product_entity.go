package entities

type ProductEntity struct {
	Id             int64
	Name           string
	Orders         []*OrderEntity        `orm:"reverse(many)"`
	ProductDetails *ProductDetailsEntity `orm:"rel(one)"`
}

func (ProductEntity) TableName() string {
	return "product"
}
