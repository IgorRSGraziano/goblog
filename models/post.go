package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	Slug    string `json:"slug"`
	User    User   `json:"user"`
}
