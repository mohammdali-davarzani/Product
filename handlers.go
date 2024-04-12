package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProducts(c echo.Context) error {
	products := []Product{}
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return err
		}
		return nil
	})

	if actionErr != nil {
		return c.JSON(http.StatusInternalServerError, actionErr.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"result": "product created successfully"})
}

func GetProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if result := db.First(&product, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	var orgProduct = new(Product)
	if result := db.First(&orgProduct, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(orgProduct).Updates(product).Error; err != nil {
			return err
		}
		return nil
	})

	if actionErr != nil {
		return c.JSON(http.StatusInternalServerError, actionErr.Error())
	}

	return c.JSON(http.StatusOK, orgProduct)
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if result := db.First(&product, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(product).Error; err != nil {
			return err
		}
		return nil
	})

	if actionErr != nil {
		return c.JSON(http.StatusInternalServerError, actionErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"result": "Product successfuly deleted."})
}
