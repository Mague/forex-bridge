package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Mague/forex-bridge/api"
	docs "github.com/Mague/forex-bridge/docs"
	"github.com/Mague/forex-bridge/utils"
)

func initalizeRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		api.Operation{}.Load(v1)
	}
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Inicio",
		})
	})
	router.POST("/signin", func(ctx *gin.Context) {
		token := utils.NewJWT("42")
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
