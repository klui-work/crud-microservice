package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	CreateProductController(c *gin.Context)
	UpdateProductService(c *gin.Context)
	DeleteProductController(c *gin.Context)
	GetProductByQueryController(c *gin.Context)
}
