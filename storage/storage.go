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
		dsn := "host=localhost user=go password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
