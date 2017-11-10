package entities

type UserGroupEntity struct {
	Id    int64
	Name  string
	Users []UserEntity `gorm:"ForeignKey:UserGroupId"`
}

func (UserGroupEntity) TableName() string {
	return "user_group"
}
