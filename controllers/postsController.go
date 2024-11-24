package controllers

import (
	"github.com/Francismensah/JSON-CRUD-API/initializers"
	"github.com/Francismensah/JSON-CRUD-API/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error creating post",
		})
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"post":    post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	// Find the Post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
