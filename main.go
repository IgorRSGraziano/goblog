package main

import (
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.Run()
}
