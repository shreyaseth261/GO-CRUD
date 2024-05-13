package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyaseth/go-crud/initializers"
	"github.com/shreyaseth/go-crud/models"
)

func PostCreate(c *gin.Context) {

	//populate data from body of req
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)
	//create a post or add to db
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) 
	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context){
	//get all the posts
	var posts [] models.Post
	initializers.DB.Find(&posts)


	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostsById(c *gin.Context){
	//get Id from query params
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post,id);

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context){
	//extract the id to update from query params
	id := c.Param("id")

	//get the data to be updated from the body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//find th post to be updated
	var post models.Post
	initializers.DB.First(&post,id);

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body,})


	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context){
	//extract the id to update from query params
	id := c.Param("id")

	//find th post to be updated
	initializers.DB.Delete(&models.Post{}, id)


	//respond with them
	c.Status(200)
}