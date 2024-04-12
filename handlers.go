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
	c.Logger().Info("GetProducts endpoint successfully worked")
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return nil
	})

	if actionErr != nil {
		c.Logger().Error(actionErr.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, actionErr.Error())
	}

	c.Logger().Infof("Prodcut %v successfully created", product)
	return c.JSON(http.StatusCreated, map[string]string{"result": "product created successfully"})
}

func GetProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if result := db.First(&product, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Logger().Error(result.Error)
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	}

	c.Logger().Infof("Product %v successfully find", product)
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if err := c.Bind(product); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	var orgProduct = new(Product)
	if result := db.First(&orgProduct, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Logger().Error(result.Error)
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(orgProduct).Updates(product).Error; err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return nil
	})

	if actionErr != nil {
		c.Logger().Error(actionErr.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, actionErr.Error())
	}

	c.Logger().Infof("product %v was updated", product)
	return c.JSON(http.StatusOK, orgProduct)
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("productID")
	product := new(Product)
	if result := db.First(&product, productId); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Logger().Error(result.Error)
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	}

	actionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(product).Error; err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return nil
	})

	if actionErr != nil {
		c.Logger().Error(actionErr.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, actionErr.Error())
	}

	c.Logger().Warnf("product %v was deleted", product)
	return c.JSON(http.StatusOK, map[string]string{"result": "Product successfuly deleted."})
}
