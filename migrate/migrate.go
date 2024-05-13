package main

import (
	"github.com/shreyaseth/go-crud/initializers"
	"github.com/shreyaseth/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}