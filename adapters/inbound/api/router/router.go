package router

import (
	"crud/adapters/inbound/api/controller"

	"github.com/gin-gonic/gin"
)

func StartProductRTL(e *gin.Engine, ctl controller.IProductController) {
	groupTest := e.Group("/api/product/health-check")
	groupTest.GET("", ctl.HealthCheck)

	groupProduct := e.Group("/api/product")
	{
		groupProduct.POST("", ctl.CreateProductController)
		groupProduct.PUT("", ctl.UpdateProductController)
		groupProduct.DELETE("", ctl.DeleteProductController)
		groupProduct.GET("", ctl.GetProductByQueryController)
	}
}
