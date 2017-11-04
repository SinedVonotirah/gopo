package entities

type ProductEntity struct {
	Id             int64
	Name           string
	Orders         []*OrderEntity        `orm:"reverse(many)"`
	ProductDetails *ProductDetailsEntity `orm:"null;rel(one);on_delete(set_null)"`
}

func (ProductEntity) TableName() string {
	return "product"
}
