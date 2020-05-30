package cache

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath = "cache.db"

	// table name
	Wallhaven = "wallhaven"
	Octodex   = "octodex"
	Polayoutu = "polayoutu"
)

var (
	sqlite = NewSQLite()
)

func NewSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db
}

func IsVisited(table, key string) (string, error) {
	query := fmt.Sprintf("SELECT value,time FROM %s WHERE key=? ORDER BY id DESC LIMIT 1", table)
	rows, err := sqlite.Query(query, key)
	if err != nil {
		return "", err
	}

	if !rows.Next() {
		rows.Close()
		return "", fmt.Errorf("key not found")
	}

	var value string
	var timestamp int64
	if err := rows.Scan(&value, &timestamp); err != nil {
		rows.Close()
		return "", err
	}
	rows.Close()

	// 超过24小时, 缓存失效
	if time.Now().Unix()-timestamp > 24*60*60 {
		if err := DeleteValue(table, key); err != nil {
			fmt.Printf("IsVisited-DeleteValue err: %s\n", err)
		}
		return "", fmt.Errorf("cache timeout-> key: %s, value: %s",
			key, value)
	}
	return value, nil
}

func SaveValue(table, key, value string) error {
	query := fmt.Sprintf("insert into %s (key,value,time) values (?,?,?)", table)
	_, err := sqlite.Exec(query, key, value, time.Now().Unix())
	return err
}

func DeleteValue(table, key string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE key=?", table)
	_, err := sqlite.Exec(query, key)
	return err
}
