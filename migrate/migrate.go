package main

import (
	"todolist/initializers"
	"todolist/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Task{})
}
