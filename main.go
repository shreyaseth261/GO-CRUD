package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyaseth/go-crud/controllers"
	"github.com/shreyaseth/go-crud/initializers"
)


func init(){
initializers.LoadEnvVariables()
initializers.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/getPostsById/:id", controllers.GetPostsById)
	r.PUT("/updatePost/:id", controllers.UpdatePost)
	r.DELETE("/deletePost/:id", controllers.DeletePost)

	r.Run() }