package persistence

import (
	"github.com/SinedVonotirah/gopo/orms/gorm/persistence/entities"
	"github.com/SinedVonotirah/gopo/shared/logging"

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

func (repo *UserRepo) GetUserById(id int64) (entities.UserEntity, error) {
	user := entities.UserEntity{Id: id}
	repo.connection.Model(&user).First(&user)
	return user, nil
}

func (repo *UserRepo) GetUserByIdWithOrders(id int64) (entities.UserEntity, error) {
	user := entities.UserEntity{Id: id}
	repo.connection.Model(&user).First(&user).Related(&user.Orders, "Orders")
	return user, nil
}
