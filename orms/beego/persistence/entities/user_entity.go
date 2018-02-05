package entities

import (
	"bytes"
	"fmt"
)

type UserEntity struct {
	Id        int64
	Name      string
	Mail      string
	UserGroup *UserGroupEntity `orm:"rel(fk)"`
	Orders    []*OrderEntity   `orm:"reverse(many)"`
}

func (UserEntity) TableName() string {
	return "user"
}

func NewUserEntity(entity *UserEntity) {
	entity.Id = 0
	entity.Name = "name"
	entity.Mail = "mail"
	entity.UserGroup = &UserGroupEntity{Id: 1, Name: "Group1"}
}

// FIXME
func (entity *UserEntity) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("User - %+v, %+v, %+v", entity.Id, entity.Name, entity.Mail))
	buffer.WriteString(entity.UserGroup.ToString())
	for _, order := range entity.Orders {
		buffer.WriteString(order.ToString())
	}
	return buffer.String()
}
