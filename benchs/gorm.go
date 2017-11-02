package benchs

import (
	"github.com/SinedVonotirah/gopo/shared/logging"

	"github.com/SinedVonotirah/gopo/persistence/db_migrator"
	"github.com/SinedVonotirah/gopo/persistence/entities"
	"github.com/SinedVonotirah/gopo/persistence/gorm"
)

const (
	connection     = "host=localhost user=postgres password=1 dbname=gopo sslmode=disable"
	migrationUrl   = "postgresql://postgres:1@localhost:5432/gopo?sslmode=disable"
	migrationsPath = "file://persistence/db_migrator/migrations/"
)

var repo *gorm.UserRepo

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GormInsert)

		connection := gorm.NewConnection(connection)
		repo = gorm.NewUserRepo(connection)
	}

}

func GormInsert(b *B) {
	var entity entities.UserEntity

	wrapExecute(b, func() {
		entities.NewUserEntity(&entity)
		db_migrator.ApplyMigrations(migrationsPath, migrationUrl, true)
	})
	for i := 0; i < b.N; i++ {
		entity.Id = 0
		err := repo.Insert(&entity)
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Gorm insert error")
			b.FailNow()
		}
	}
}
