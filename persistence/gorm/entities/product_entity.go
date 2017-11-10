package entities

type ProductEntity struct {
	Id               int64
	Name             string
	ProductDetailsId int64
	ProductDetails   ProductDetailsEntity `gorm:"ForeignKey:ProductDetailsId"`
}

func (ProductEntity) TableName() string {
	return "product"
}
