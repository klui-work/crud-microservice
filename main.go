package main

import (
	"crud/infra/conf"
	"fmt"

	// "net/http"
	// "github.com/gin-gonic/gin"

	"crud/infra/persistence"
)

func main() {
	// r := gin.Default()
	conf.LoadConfig()
	mongo := persistence.SetupMongoDB()
	fmt.Println(mongo)
	// r.GET("/health-check", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "statusOK",
	// 	})
	// })
	// r.Run(conf.Env.Port)

}
