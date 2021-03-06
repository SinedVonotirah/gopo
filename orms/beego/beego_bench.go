package beego

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"

	"github.com/SinedVonotirah/gopo/shared/bench"
	"github.com/SinedVonotirah/gopo/orms/beego/persistence"
	"github.com/SinedVonotirah/gopo/orms/beego/persistence/entities"

	"github.com/SinedVonotirah/gopo/orms/beego/persistence/repo"
)

var beegoUserRepo *repo.UserRepo
var beegoOrderRepo *repo.OrderRepo

func init() {
	st := bench.NewSuite("beego")
	st.InitF = func() {
		st.AddBenchmark("GetUserById", 1*bench.ORM_MULTI, BeegoGetUserById)
		st.AddBenchmark("GetUserByIdWithOrders", 1*bench.ORM_MULTI, BeegoGetUserByIdWithOrders)

		connection := persistence.NewConnection(bench.ConnectionStr)
		beegoUserRepo = repo.NewUserRepo(connection)
		beegoOrderRepo = repo.NewOrderRepo(connection)
	}
}

func BeegoGetUserById(b *bench.B) {
	for i := 0; i < b.N; i++ {
		_, err := beegoUserRepo.GetUserById(1)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("BeegoGetUserById bench error")
			b.FailNow()
		}
	}
}

func BeegoGetUserByIdWithOrders(b *bench.B) {
	for i := 0; i < b.N; i++ {
		_, err := beegoUserRepo.GetUserByIdWithOrders(1)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("BeegoGetUserByIdWithOrders bench error")
			b.FailNow()
		}
	}
}

func BeegoCreateUser(b *bench.B) {
	var entity entities.UserEntity

	for i := 0; i < b.N; i++ {
		bench.WrapExecute(b, func() {
			entities.NewUserEntity(&entity)
		})
		id, err := beegoUserRepo.CreateUser(&entity)
		fmt.Println(id)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("CreateUser bench error")
			b.FailNow()
		}
	}
}

func BeegoCreateOrder(b *bench.B) {
	var entity entities.OrderEntity

	for i := 0; i < b.N; i++ {
		bench.WrapExecute(b, func() {
			entities.NewOrderEntity(&entity)
		})
		id, err := beegoOrderRepo.CreateOrder(&entity)
		fmt.Println(id)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("CreateUser bench error")
			b.FailNow()
		}
	}
}

func GetOrderByUserName(b *bench.B) {
	for i := 0; i < b.N; i++ {
		entities, err := beegoOrderRepo.GetOrdersByUserName("User1")
		fmt.Println(entities)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GetOrderByUserName bench error")
			b.FailNow()
		}
	}
}

func GetOrdersByProductName(b *bench.B) {
	for i := 0; i < b.N; i++ {
		entities, err := beegoOrderRepo.GetOrdersByProductName("Product1")
		fmt.Println(entities)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GetOrderByProductName bench error")
			b.FailNow()
		}
	}
}

func GetOrderById(b *bench.B) {
	for i := 0; i < b.N; i++ {
		entity, err := beegoOrderRepo.GetOrderById(1)
		fmt.Println(&entity)
		fmt.Println(entity.User.UserGroup)
		fmt.Println(entity.Products[0].ProductDetails)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GetOrderById bench error")
			b.FailNow()
		}
	}
}

func GetUsersByGroupName(b *bench.B) {
	for i := 0; i < b.N; i++ {
		entities, err := beegoUserRepo.GetUsersByGroupName("Group1")
		fmt.Println(entities)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("GetOrderByProductName bench error")
			b.FailNow()
		}
	}
}
