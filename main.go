package main

import (
	"goblog/controllers"
	"goblog/models"

	_ "goblog/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						token
func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/user/login", controllers.Login)
	r.Run()
}
