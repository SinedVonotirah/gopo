package benchs

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"

	"github.com/SinedVonotirah/gopo/persistence/db_migrator"
	"github.com/SinedVonotirah/gopo/persistence/entities"
	"github.com/SinedVonotirah/gopo/persistence/gorm"
)

var userRepo *gorm.UserRepo
var orderRepo *gorm.OrderRepo

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		//		st.AddBenchmark("Insert", 2000*ORM_MULTI, GormInsertOrder)
		st.AddBenchmark("Insert", 1*ORM_MULTI, GormGetById)

		connection := gorm.NewConnection(connection)
		userRepo = gorm.NewUserRepo(connection)
		orderRepo = gorm.NewOrderRepo(connection)
	}

}

func GormInsertUser(b *B) {
	var entity entities.UserEntity

	wrapExecute(b, func() {
		entities.NewUserEntity(&entity)
		db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		entity.Id = 0
		err := userRepo.Insert(&entity)
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
		err := orderRepo.Insert(&entity)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}

func GormGetById(b *B) {

	wrapExecute(b, func() {
		db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		usr, err := userRepo.GetUserById(1)
		fmt.Println(usr)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}
