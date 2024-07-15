package main

import (
	"goblog/controllers"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/user/login", controllers.Login)
	r.Run()
}
