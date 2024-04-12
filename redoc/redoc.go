package redoc

import (
	"encoding/json"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

func RedocTemplate(c echo.Context) error {
	// Retrieve endpoint information from your Go code

	// Serve the OpenAPI JSON
	return c.Render(http.StatusOK, "redoc.html", map[string]string{})
}

func Endpoints(c echo.Context) error {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Your API",
			Description: "Description of your API",
			Version:     "1.0.0",
		},
		Paths: openapi3.NewPaths(),
	}

	productsPathItem := &openapi3.PathItem{}
	productPathItem := &openapi3.PathItem{}

	getProducts := &openapi3.Operation{
		OperationID: "Get all products",
	}
	createProduct := &openapi3.Operation{
		OperationID: "Add new product",
	}
	getProduct := &openapi3.Operation{
		OperationID: "Get product by id",
	}
	updateProduct := &openapi3.Operation{
		OperationID: "Update product by id",
	}
	deleteProduct := &openapi3.Operation{
		OperationID: "Delete product by id",
	}

	productsPathItem.Get = getProducts
	productsPathItem.Post = createProduct
	productPathItem.Get = getProduct
	productPathItem.Put = updateProduct
	productPathItem.Delete = deleteProduct

	// Add the createProductPathItem to the Paths field of the openapi3.T instance
	doc.Paths.Set("/api/v1/products", productsPathItem)
	doc.Paths.Set("/api/v1/product/:productID", productPathItem)

	// Marshal OpenAPI document to JSON
	openAPIJSON, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	// Serve the OpenAPI JSON
	return c.JSONBlob(http.StatusOK, openAPIJSON)
	// Retrieve endpoint information from your Go code

}
