package middleware

import (
	"errors"
	"goblog/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func AuthMiddleware(c *gin.Context) {

	token, err := extractBearerToken(c.GetHeader("token"))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user, err := services.FindUserByToken(token)

	if user == nil || err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("user", user)

	c.Next()

}
