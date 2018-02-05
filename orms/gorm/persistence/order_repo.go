package persistence

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/SinedVonotirah/gopo/orms/gorm/persistence/entities"

	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	connection *gorm.DB
}

func NewOrderRepo(connection *Connection) *OrderRepo {
	return &OrderRepo{connection.gorm}
}

func (repo *OrderRepo) Insert(entity *entities.OrderEntity) error {
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
