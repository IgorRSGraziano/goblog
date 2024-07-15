package controllers

import (
	_ "goblog/models"
	"goblog/services"

	"github.com/gin-gonic/gin"
)

// Login godoc
//
//	@Summary		User login
//	@Description	User authentication, to get a token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			input	body		services.LoginInput	true	"User login payload"
//	@Success		200		{object}	object{token=string, user=models.User}
//	@Failure		401		{object}	object{error=string}
//	@Failure		500		{object}	object{error=string}
//	@Router			/user/login [post]
func Login(c *gin.Context) {
	var input services.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	token, user, err := services.Login(input)

	if err != nil || token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"token": token, "user": user})
}
