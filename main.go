package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

var db = GetConnection()

type Endpoint struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	InitializeDB(db)
	e := echo.New()
	e.Renderer = &TemplateRender{
		templates: template.Must(template.ParseGlob("./*.html")),
	}
	e.GET("/openapi", Endpoints)
	e.GET("/docs", RedocTemplate)
	g := e.Group("/api/v1")
	g.GET("/products", GetProducts)
	g.POST("/products", CreateProduct)
	g.GET("/product/:productID", GetProduct)
	g.PUT("/product/:productID", UpdateProduct)
	g.DELETE("/product/:productID", DeleteProduct)
	e.Start(":8000")
}
