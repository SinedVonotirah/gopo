package benchs

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"

	"github.com/SinedVonotirah/gopo/persistence/beego"
)

var beegoUserRepo *beego.UserRepo

func init() {
	st := NewSuite("beego")
	st.InitF = func() {
		st.AddBenchmark("GetById", 1*ORM_MULTI, GetById)

		connection := beego.NewConnection(connection)
		beegoUserRepo = beego.NewUserRepo(connection)
	}

}

func GetById(b *B) {

	wrapExecute(b, func() {
		//db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		entity, err := beegoUserRepo.GetUserById(1)
		fmt.Println(entity)
		for _, order := range entity.Orders {
			fmt.Println(order)
			for _, product := range order.Product {
				fmt.Println(product)
				fmt.Println(product.ProductDetails)
			}
		}
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}
