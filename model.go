package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func InitializeDB(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}
