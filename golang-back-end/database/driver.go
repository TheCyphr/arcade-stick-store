package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbConn.SQL = db
	return dbConn, nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
