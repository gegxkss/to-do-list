package controllers

import (
	"todolist/initializers"
	"todolist/models"
	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var authInput models.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}
	if authInput.Username == "" || authInput.Password == "" {
		c.JSON(400, gin.H{
			"message": "Username or Email or Password are not filled",
		})
		return
	}
	user := models.User{
		Username: authInput.Username,
		Email:    authInput.Email,
		Password: authInput.Password,
	}
	initializers.DB.Create(&user)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
