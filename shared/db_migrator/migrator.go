package db_migrator

import (
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	"github.com/SinedVonotirah/gopo/shared/logging"
)

func ApplyMigrations(folder string, connectionString string, drop bool) {
	m, err := migrate.New(folder, connectionString)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Creating migration error")
	}
	if drop {
		m.Drop()
	}
	m.Up()
}
