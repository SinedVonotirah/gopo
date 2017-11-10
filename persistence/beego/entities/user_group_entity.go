package entities

type UserGroupEntity struct {
	Id    int64
	Name  string
	Users []*UserEntity `orm:"reverse(many)"`
}

func (UserGroupEntity) TableName() string {
	return "user_group"
}
