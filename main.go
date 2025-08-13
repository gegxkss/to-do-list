package main

import (
	"todolist/controllers"
	"todolist/initializers"
	"todolist/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Registration)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.LogOut)
	r.GET("/:username/profile", middlewares.CheckAuth, controllers.UserProfile)
	r.POST("/:username/addTask", middlewares.CheckAuth, controllers.AddTask)
	r.Run(":8080")
}
