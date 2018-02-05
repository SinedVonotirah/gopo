package entities

import (
	"bytes"
	"fmt"
)

type UserGroupEntity struct {
	Id    int64
	Name  string
	Users []*UserEntity `orm:"reverse(many)"`
}

func (UserGroupEntity) TableName() string {
	return "user_group"
}

func (entity *UserGroupEntity) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("UserGroup - %+v, %+v", entity.Id, entity.Name))
	for _, user := range entity.Users {
		buffer.WriteString(user.ToString())
	}
	return buffer.String()
}
