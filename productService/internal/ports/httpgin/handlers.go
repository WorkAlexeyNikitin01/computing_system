package service

import (
	"lab/productService/internal/app"
	"lab/productService/internal/product"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allListProducts(a app.AppProductInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody productRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error get list products", err)
			c.JSON(400, productErr(err))
			return
		}

		products, err := a.AllListProducts()
		if err != nil {
			c.JSON(200, productErr(err))
			log.Println("error get list products")
			return
		}
		log.Println("Success list products", http.StatusOK, "len products", len(products))
		c.Status(http.StatusOK)
		productsSuccess(products)
		c.JSON(200, products[0])
		log.Default()
	}
}

func createProduct(a app.AppProductInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody productRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error create product", err)
			c.JSON(400, productErr(err))
			return
		}

		p, err := a.CreateProduct(product.Product{
			Code: reqBody.code,
			Name: reqBody.name,
			Description: reqBody.description,
			Price: reqBody.price,
		})
		if err != nil {
			c.JSON(200, productErr(err))
			log.Println("error create product")
			return
		}
		log.Println("Success product", http.StatusOK, "id product", p.Id)
		c.Status(http.StatusOK)
		c.JSON(200, productSuccess(p))
		log.Default()
	}
}