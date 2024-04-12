package main

import (
	"github.com/labstack/echo/v4"
)

var db = GetConnection()

func main() {
	InitializeDB(db)
	e := echo.New()
	g := e.Group("/api/v1")
	g.GET("/products", GetProducts)
	g.POST("/products", CreateProduct)
	g.GET("/product/:productID", GetProduct)
	g.PUT("/product/:productID", UpdateProduct)
	g.DELETE("/product/:productID", DeleteProduct)
	e.Start(":8000")
}
