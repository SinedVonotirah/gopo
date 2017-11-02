package gorm

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/SinedVonotirah/gopo/persistence/entities"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection *Connection) *UserRepo {
	return &UserRepo{connection.gorm}
}

func (repo *UserRepo) Insert(entity *entities.UserEntity) error {
	db := repo.connection.Begin()
	err := db.Create(entity).Error
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during serving")
		db.Rollback()
		return err
	} else {
		db.Commit()
	}
	return nil
}
