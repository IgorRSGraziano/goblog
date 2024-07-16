package services

import (
	"crypto/md5"
	"encoding/hex"
	"goblog/models"

	"github.com/golang-jwt/jwt/v5"
)

// just for testing purposes :p
var jwtSecretKey []byte = []byte("b0276b14b4a27a9286465d35c357a06a52845a1e7cd9d0f5a96b7ae05e010e0c")

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateJwt(user *models.User) (string, error) {

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

func FindUserByToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	claimedID, ok := claims["id"].(float64)

	if !ok {
		return nil, err
	}

	user, err := FindUserById(uint(claimedID))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserById(id uint) (*models.User, error) {
	var user *models.User
	err := models.DB.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
