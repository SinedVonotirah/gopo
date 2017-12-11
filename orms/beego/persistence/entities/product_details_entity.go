package entities

type ProductDetailsEntity struct {
	Id          int64
	Description string
	Product     *ProductEntity `orm:"reverse(one)"`
}

func (ProductDetailsEntity) TableName() string {
	return "product_details"
}
