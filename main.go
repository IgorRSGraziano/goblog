package main

import (
	"goblog/controllers"
	"goblog/middleware"
	"goblog/models"

	_ "goblog/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						token
func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/user/login", controllers.Login)
	r.Use(middleware.AuthMiddleware).GET("/user", controllers.GetLoggedInUser)

	r.Run()
}
