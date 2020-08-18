package db

import (
	"database/sql"

	"github.com/astaxie/beego/logs"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath = "database.db"
)

var (
	sqliteDB *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	sqliteDB = db
}

func Get() *sql.DB {
	return sqliteDB
}

func Close() {
	err := sqliteDB.Close()
	if err != nil {
		logs.Error("%s", err)
	}
}
