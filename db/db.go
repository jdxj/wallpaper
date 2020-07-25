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
	openDB()
}

func openDB() {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	sqliteDB = db
	logs.Info("open database success")
}

func Close() {
	err := sqliteDB.Close()
	if err != nil {
		logs.Error("%s", err)
	}
}
