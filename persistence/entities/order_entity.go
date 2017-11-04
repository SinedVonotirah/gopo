package entities

type OrderEntity struct {
	Id      int64
	Name    string
	User    *UserEntity      `orm:"rel(fk)"`
	Product []*ProductEntity `orm:"rel(m2m)"`
}

func (OrderEntity) TableName() string {
	return "order"
}
