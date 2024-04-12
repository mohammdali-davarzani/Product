package database

import (
	"errors"
	"fmt"
	"os"
	"shop/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var doOnce sync.Once
var singleton *gorm.DB

func Connect() error {
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
	db.AutoMigrate(&models.Product{})

	return err
}

func GetConnection() (*gorm.DB, error) {
	if singleton == nil {
		err := Connect()
		if err != nil {
			return nil, errors.New("database connection is not initialized")
		}
	}
	return singleton, nil
}
