package xorm

import (
	"github.com/SinedVonotirah/gopo/persistence/entities"
	"github.com/go-xorm/xorm"
)

type UserRepo struct {
	connection *xorm.Engine
}

func NewUserRepo(connection *Connection) *UserRepo {
	return &UserRepo{connection.xorm}
}

func (repo *UserRepo) GetUserById(id int64) (entities.UserEntity, error) {
	entity := entities.UserEntity{}
	_, err := repo.connection.Id(id).Get(&entity)
	return entity, err
}
