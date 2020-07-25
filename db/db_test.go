package db

import "testing"

func TestOpenDB(t *testing.T) {
	if err := sqliteDB.Ping(); err != nil {
		t.Fatalf("%s\n", err)
	}
	sqliteDB.Close()
}
