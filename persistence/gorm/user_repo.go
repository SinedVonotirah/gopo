package gorm

import "github.com/jinzhu/gorm"

type UserRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection Connection) *UserRepo {
	return &UserRepo{connection.gorm}
}
