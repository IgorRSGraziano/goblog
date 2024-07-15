package services

import "goblog/models"

func FindAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := models.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(post models.Post, userID uint) error {
	post.UserID = userID

	var slug string
	for _, char := range post.Title {
		if char == ' ' {
			slug += "-"
		} else {
			slug += string(char)
		}

	}

	post.Slug = slug

	err := models.DB.Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func FindPostById(id string) (models.Post, error) {
	var post models.Post
	err := models.DB.First(&post, id).Error
	if err != nil {
		return post, err
	}
	return post, nil
}
