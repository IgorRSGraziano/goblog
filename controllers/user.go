package controllers

import (
	"goblog/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input services.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(input)

	if err != nil || token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
