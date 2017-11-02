package entities

type ProductDetailsEntity struct {
	ProductId   string
	Description string
}

func (ProductDetailsEntity) TableName() string {
	return "product_details"
}
