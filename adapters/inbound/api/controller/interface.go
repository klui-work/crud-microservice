package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	HealthCheck(c *gin.Context)
	CreateProductController(c *gin.Context)
	UpdateProductController(c *gin.Context)
	DeleteProductController(c *gin.Context)
	GetProductByQueryController(c *gin.Context)
}
