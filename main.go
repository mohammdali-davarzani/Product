package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"shop/database"
	"shop/handlers"
	"shop/redoc"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func WithDBConnection(handlerFunc func(c echo.Context, db *gorm.DB) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		dbConn, err := database.GetConnection()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadGateway, "Can't connect to Database")
		}
		return handlerFunc(c, dbConn)

	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("cannot load .env file")
	}

	err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connecting to database")
		return
	}

	e := echo.New()
	e.Renderer = &TemplateRender{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.GET("/openapi", redoc.Endpoints)
	e.GET("/docs", redoc.RedocTemplate)
	g := e.Group("/api/v1")
	g.GET("/products", WithDBConnection(handlers.GetProducts))
	g.POST("/products", WithDBConnection(handlers.CreateProduct))
	g.GET("/product/:productID", WithDBConnection(handlers.GetProduct))
	g.PUT("/product/:productID", WithDBConnection(handlers.UpdateProduct))
	g.DELETE("/product/:productID", WithDBConnection(handlers.DeleteProduct))
	log.Fatal(e.Start(":8000"))
}
