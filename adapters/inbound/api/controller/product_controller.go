package controller

import (
	"crud/application/dtos"
	"crud/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func (p ProductController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Product API is running",
	})
}

func (p ProductController) CreateProductController(c *gin.Context) {
	var product dtos.ProductDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := p.productService.CreateProductService(product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "create product successfully",
	})
}

func (p ProductController) UpdateProductController(c *gin.Context) {
	var product dtos.ProductDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := p.productService.UpdateProductService(product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update product successfully",
	})

}

func (p ProductController) DeleteProductController(c *gin.Context) {
	var productId string
	if err := c.ShouldBindJSON(productId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := p.productService.DeleteProductService(productId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "delete product successfully",
	})
}

func (p ProductController) GetProductByQueryController(c *gin.Context) {
	var query dtos.QueryProduct
	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	results, err := p.productService.GetProductByQueryService(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": results,
	})
}

func NewProductController(productService service.IProductService) IProductController {
	return &ProductController{
		productService: productService,
	}
}
