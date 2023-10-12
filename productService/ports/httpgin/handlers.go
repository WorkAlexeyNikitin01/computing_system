package httpgin

import (
	"lab/productService/app"
	"lab/productService/product"

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
			log.Println("error create product reqBody", err)
			c.JSON(400, productErr(err))
			return
		}
		log.Println(reqBody.Code, reqBody.Name, reqBody.Description, reqBody.Price)
		p, err := a.CreateProduct(product.Product{
			Code: reqBody.Code,
			Name: reqBody.Name,
			Description: reqBody.Description,
			Price: reqBody.Price,
		})
		if err != nil {
			c.JSON(200, productErr(err))
			log.Println("error create product app")
			return
		}
		log.Println("Success product", http.StatusOK, "id product", p.Id)
		c.Status(http.StatusOK)
		c.JSON(200, productSuccess(p))
		log.Default()
	}
}

func deleteProduct(a app.AppProductInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req productRequest
		err := c.Bind(&req)
		if err != nil {
			log.Fatal(err)
		}
		result, err := a.DeleteProduct(req.Code)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, productSuccess(result))
		log.Println("delete product", result.Id)
	}
}

func getProduct(a app.AppProductInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody productRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error get product bind", err)
			c.JSON(400, productErr(err))
			return
		}

		product, err := a.GetProduct(reqBody.Code)
		if err != nil {
			c.JSON(200, productErr(err))
			log.Println("error get product")
			return
		}
		log.Println("Success product", http.StatusOK, "product", product)
		c.Status(http.StatusOK)
		productSuccess(product)
		c.JSON(200, product)
		log.Default()
	}
}