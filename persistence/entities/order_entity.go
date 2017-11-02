package entities

type OrderEntity struct {
	Id   string
	Name string
	User UserEntity
}

func (OrderEntity) TableName() string {
	return "order"
}
