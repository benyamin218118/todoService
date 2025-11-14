package db

import (
	"database/sql"
	"time"

	"github.com/benyamin218118/todoService/domain"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(conf *domain.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", conf.DBDSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(32)           // max open connections
	db.SetMaxIdleConns(8)            // max idle connections
	db.SetConnMaxLifetime(time.Hour) // how long to keep a connection alive

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
