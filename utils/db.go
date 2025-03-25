package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"sync"
)

type DB struct {
	*sql.DB
}

func createDB() *DB {
	dbAddress := os.Getenv("DB_ADDRESS")
	if dbAddress == "" {
		panic("invalid database address")
	}
	slog.Info("Connecting to database " + dbAddress)
	db, err := sql.Open("postgres", dbAddress)
	if err != nil {
		slog.Error("Fail to connect to db " + err.Error())
		return nil
	}
	slog.Info("Connected to database success!")
	return &DB{db}
}

var dbOnce sync.Once
var appDB *DB

func GetDB() *DB {
	dbOnce.Do(func() {
		appDB = createDB()
	})
	return appDB
}
