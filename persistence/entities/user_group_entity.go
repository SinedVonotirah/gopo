package entities

type UserGroupEntity struct {
	Id    int64
	Name  string
	Users []UserEntity
}

func (UserGroupEntity) TableName() string {
	return "user_group"
}
