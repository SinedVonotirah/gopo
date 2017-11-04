package benchs

import (
	"fmt"

	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/SinedVonotirah/gopo/persistence/xorm"
)

var xormUserRepo *xorm.UserRepo

func init() {
	st := NewSuite("xorm")
	st.InitF = func() {
		st.AddBenchmark("GetById", 1*ORM_MULTI, XormGetById)

		connection := xorm.NewConnection(connection)
		xormUserRepo = xorm.NewUserRepo(connection)
	}

}

func XormGetById(b *B) {

	wrapExecute(b, func() {
		//db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		entity, err := xormUserRepo.GetUserById(1)
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
			}).Error("Xorm get error")
			b.FailNow()
		}
	}
}
