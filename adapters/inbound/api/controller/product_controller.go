package controller

import (
	"crud/application/dtos"
	"crud/domain/service"
	"net/http"

	"crud/adapters/inbound/api/validation"

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

	// bind json
	var product dtos.ProductDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validation
	if err := validation.ValidationCreateProduct(product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create product
	err := p.productService.CreateProductService(product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "create product successfully",
	})
}

func (p ProductController) UpdateProductController(c *gin.Context) {

	// bind json
	var product dtos.ProducUpdatetDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validation
	if err := validation.ValidationUpdateProduct(product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update product
	if err := p.productService.UpdateProductService(product); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update product successfully",
	})

}

func (p ProductController) DeleteProductController(c *gin.Context) {

	// get product id from param
	productId := c.Param("product_id")

	// validation
	if err := validation.ValidationDeleteProduct(productId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// delete product
	if err := p.productService.DeleteProductService(productId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "delete product successfully",
	})
}

func (p ProductController) GetProductByQueryController(c *gin.Context) {

	// query
	query := dtos.QueryProduct{
		ProductId:   c.Query("product_id"),
		ProductName: c.Query("product_name"),
	}

	results, err := p.productService.GetProductByQueryService(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
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
