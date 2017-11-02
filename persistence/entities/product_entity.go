package entities

type ProductEntity struct {
	Id             string
	Name           string
	ProductDetails ProductDetailsEntity
}

func (ProductEntity) TableName() string {
	return "product"
}
