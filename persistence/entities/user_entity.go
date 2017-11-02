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

func NewUserEntity(entity *UserEntity) {
	entity.Id = 0
	entity.Name = "name"
	entity.Mail = "mail"
}
