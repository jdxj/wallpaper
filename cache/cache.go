package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v2"

	_ "github.com/mattn/go-sqlite3"
)

type storeValue struct {
	Timestamp int64  `json:"timestamp"`
	Data      []byte `json:"data"`
}

const (
	dbPath = "cache.db"
)

var (
	DB = NewBadger()
)

func NewBadger() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		panic(err)
	}
	return db
}

func IsVisited(key []byte) ([]byte, error) {
	var result []byte
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		result, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		return nil, err
	}

	// 检查是否超时
	sv := &storeValue{}
	if err := json.Unmarshal(result, sv); err != nil {
		return nil, err
	}

	// 超过24小时, 缓存失效
	if time.Now().Unix()-sv.Timestamp > 24*60*60 {
		if err := DeleteValue(key); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("this key timeout: %s", key)
	}
	return sv.Data, nil
}

func DeleteValue(key []byte) error {
	err := DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	return err
}

func SaveValue(key, value []byte) error {
	sv := &storeValue{
		Timestamp: time.Now().Unix(),
		Data:      value,
	}
	result, err := json.Marshal(sv)
	if err != nil {
		return err
	}

	err = DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, result)
	})
	return err
}
