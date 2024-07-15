package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"goblog/models"
)

func main() {
	models.ConnectDatabase()

	// Create admin user

	digestPassword := md5.Sum([]byte("admin"))

	adminEmail := "admin@admin.com"

	var user models.User
	models.DB.First(user, "email = ?", adminEmail)

	if user.ID == 0 {
		models.DB.Create(&models.User{
			Name:     "admin",
			Email:    adminEmail,
			Password: hex.EncodeToString(digestPassword[:]),
		})
		fmt.Println("Admin user created")
	}

}
