package entities

type OrderEntity struct {
	Id      int64
	Name    string
	UserId  int64
	Product []ProductEntity `gorm:"many2many:order_products;"`
}

func (OrderEntity) TableName() string {
	return "order"
}

