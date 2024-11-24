package main

import (
	"github.com/Francismensah/JSON-CRUD-API/initializers"
	"github.com/Francismensah/JSON-CRUD-API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
