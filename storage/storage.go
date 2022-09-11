package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		dsn := "host=localhost user=go password=gorm dbname=postgres port=6432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can´t no open DB %v", err)
		}
		fmt.Println("Conectado")
	})
}

// return unique instances of db
func DB() *gorm.DB {
	return db
}
