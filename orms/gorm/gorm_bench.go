package gorm

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"

	"github.com/SinedVonotirah/gopo/shared/db_migrator"
	"github.com/SinedVonotirah/gopo/orms/gorm/persistence/repo"
	"github.com/SinedVonotirah/gopo/orms/gorm/persistence"
	"github.com/SinedVonotirah/gopo/orms/gorm/persistence/entities"
	"github.com/SinedVonotirah/gopo/shared/bench"
)

var gormUserRepo *repo.UserRepo
var gormOrderRepo *repo.OrderRepo

func init() {
	st := bench.NewSuite("gorm")
	st.InitF = func() {

		st.AddBenchmark("GetUserById", 1*bench.ORM_MULTI, GormGetUserById)
		st.AddBenchmark("GetUserByIdWithOrders", 1*bench.ORM_MULTI, GormGetUserByIdWithOrders)

		connection := persistence.NewConnection(bench.ConnectionStr)
		gormUserRepo = repo.NewUserRepo(connection)
		gormOrderRepo = repo.NewOrderRepo(connection)
	}

}

func GormGetUserById(b *bench.B) {
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

func GormGetUserByIdWithOrders(b *bench.B) {
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

func GormCreateUser(b *bench.B) {
	var entity entities.UserEntity

	for i := 0; i < b.N; i++ {
		bench.WrapExecute(b, func() {
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

func GormInsertOrder(b *bench.B) {
	var entity entities.OrderEntity

	bench.WrapExecute(b, func() {
		//entity = entities.NewOrder()
		//fmt.Println(entity)
		//db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
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
