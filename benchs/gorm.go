package benchs

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"

	"github.com/SinedVonotirah/gopo/persistence/db_migrator"
	"github.com/SinedVonotirah/gopo/persistence/gorm"
	"github.com/SinedVonotirah/gopo/persistence/gorm/entities"
)

var gormUserRepo *gorm.UserRepo
var gormOrderRepo *gorm.OrderRepo

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {

		st.AddBenchmark("GetUserById", 1*ORM_MULTI, GormGetUserById)
		st.AddBenchmark("GetUserByIdWithOrders", 1*ORM_MULTI, GormGetUserByIdWithOrders)

		connection := gorm.NewConnection(connectionStr)
		gormUserRepo = gorm.NewUserRepo(connection)
		gormOrderRepo = gorm.NewOrderRepo(connection)
	}

}

func GormGetUserById(b *B) {
	for i := 0; i < b.N; i++ {
		_, err := gormUserRepo.GetUserById(1)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GormGetUserById error")
			b.FailNow()
		}
	}
}

func GormGetUserByIdWithOrders(b *B) {
	for i := 0; i < b.N; i++ {
		_, err := gormUserRepo.GetUserByIdWithOrders(1)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GormGetUserByIdWithOrders error")
			b.FailNow()
		}
	}
}

func GormCreateUser(b *B) {
	var entity entities.UserEntity

	for i := 0; i < b.N; i++ {
		wrapExecute(b, func() {
			entities.NewUserEntity(&entity)
		})
		entity.Id = 0
		err := gormUserRepo.Insert(&entity)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}

func GormInsertOrder(b *B) {
	var entity entities.OrderEntity

	wrapExecute(b, func() {
		//entity = entities.NewOrder()
		fmt.Println(entity)
		db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		err := gormOrderRepo.Insert(&entity)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}
