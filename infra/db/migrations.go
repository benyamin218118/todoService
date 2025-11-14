package db

import (
	"github.com/benyamin218118/todoService/domain"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(conf *domain.Config) error {
	dbConn, err := GetConnection(conf)
	if err != nil {
		panic(err)
	}
	driver, _ := mysql.WithInstance(dbConn, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./infra/db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
