package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func New() {
	newPostgresDB()
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		dsn := "host=localhost user=ramon password=GormPass dbname=goland port=6432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can´t no open DB %v", err)
		}
		fmt.Println("Connect ok")
	})
}

// return unique instances of db
func DB() *gorm.DB {
	return db
}
