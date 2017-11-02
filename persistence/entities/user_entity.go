package entities

type UserEntity struct {
	Id        int64
	Name      string
	Mail      string
	UserGroup UserGroupEntity
}

func (UserEntity) TableName() string {
	return "user"
}
