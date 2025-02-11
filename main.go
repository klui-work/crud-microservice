package main

import (
	"crud/infra/conf"

	"net/http"

	"crud/adapters/inbound/api/router"
	"crud/infra/injector"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	conf.LoadConfig()
	initAdapter := injector.InitializeAdapter()
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "statusOK",
		})
	})
	router.StartProductRTL(r, initAdapter.InitialProductController)
	r.Run(conf.Env.Port)

}
