package utils

import (
	"database/sql"

	"project-name/internal/database"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(dsn string) error {
	db, err := database.New(dsn)
	if err != nil {
		return err
	}

	err = migrateSchema(db.GetConn())
	if err != nil {
		return err
	}

	return nil
}

func migrateSchema(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MultiStatementEnabled: false,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migration", "postgres", driver)
	if err != nil {
		return err
	}

	return m.Up()
}
