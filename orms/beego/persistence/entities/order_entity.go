package entities

type OrderEntity struct {
	Id       int64
	Name     string
	User     *UserEntity      `orm:"rel(fk)"`
	Products []*ProductEntity `orm:"rel(m2m)"`
}

func (OrderEntity) TableName() string {
	return "order"
}

func NewOrderEntity(entity *OrderEntity) {
	entity.Id = 0
	entity.Name = "Order"
	entity.User = &UserEntity{Id: 1, Name: "name", UserGroup: &UserGroupEntity{Id: 1, Name: "Group1"}}
	entity.Products = []*ProductEntity{
		&ProductEntity{Id: 1, Name: "product", ProductDetails: &ProductDetailsEntity{Id: 1, Description: "desc"}},
	}
}
