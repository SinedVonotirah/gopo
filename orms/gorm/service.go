package gorm

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/SinedVonotirah/gopo/orms/gorm/persistence"
	"github.com/SinedVonotirah/gopo/orms/gorm/persistence/entities"
	"github.com/SinedVonotirah/gopo/shared/bench"
)

var gormUserRepo *persistence.UserRepo
var gormOrderRepo *persistence.OrderRepo

/*
func init() {
	st := bench.NewSuite("gorm")
	st.InitF = func() {

		st.AddBenchmark("GetUserById", 1*ORM_MULTI, GormGetUserById)
		st.AddBenchmark("GetUserByIdWithOrders", 1*ORM_MULTI, GormGetUserByIdWithOrders)

		connection := persistence.NewConnection(connectionStr)
		gormUserRepo = persistence.NewUserRepo(connection)
		gormOrderRepo = persistence.NewOrderRepo(connection)
	}

}
*/

func InitBenchSuite(connectionUrl string, repeatsCount int) {
	st := bench.NewSuite("gorm")
	st.InitF = func() {

		st.AddBenchmark("GetUserById", repeatsCount, GetUserById)
		st.AddBenchmark("GetUserByIdWithOrders", repeatsCount, GetUserByIdWithOrders)

		connection := persistence.NewConnection(connectionUrl)
		gormUserRepo = persistence.NewUserRepo(connection)
		gormOrderRepo = persistence.NewOrderRepo(connection)
	}

}

func GetUserById(b *bench.B) {
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

func GetUserByIdWithOrders(b *bench.B) {
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

/*
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
*/
