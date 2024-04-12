package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var doOnce sync.Once
var singleton *gorm.DB

func GetConnection() *gorm.DB {
	doOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			fmt.Println("cannot load .env file")
		}

		var (
			user     = os.Getenv("DB_USER")
			password = os.Getenv("DB_PASSWORD")
			host     = os.Getenv("DB_HOST")
			port     = os.Getenv("DB_PORT")
			dbname   = os.Getenv("DB_NAME")
		)

		connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			user,
			password,
			host,
			port,
			dbname,
		)

		db, err := gorm.Open(
			postgres.Open(connURL),
			&gorm.Config{},
		)
		if err != nil {
			panic(err)
		}
		singleton = db
	})
	return singleton
}
