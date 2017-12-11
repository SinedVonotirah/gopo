package entities

type ProductDetailsEntity struct {
	Id          int64
	Description string
}

func (ProductDetailsEntity) TableName() string {
	return "product_details"
}
