package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"goblog/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey []byte

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func initJwt() {
	if jwtSecretKey != nil {
		return
	}

	if env := os.Getenv("JWT_SECRET_KEY"); env != "" {
		jwtSecretKey = []byte(env)
	} else {
		jwtSecretKey = generateRandomBytes(32)
	}
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateJwt(user *models.User) (string, error) {
	initJwt()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(input LoginInput) (string, *models.User, error) {
	var user *models.User

	digestPassword := md5.Sum([]byte(input.Password))

	models.DB.Where("email = ? AND password = ?", input.Email, hex.EncodeToString(digestPassword[:])).First(&user)

	if user.ID == 0 {
		return "", user, nil
	}

	token, err := generateJwt(user)

	if err != nil {
		return "", user, err
	}

	return token, user, nil
}
