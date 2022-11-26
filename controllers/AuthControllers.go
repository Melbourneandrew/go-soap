package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/melbourneandrew/go-soap/m/v2/models"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
	}

	fmt.Println(input.Username)
	fmt.Println(input.Password)

	c.JSON(200, gin.H{"token": "big daddy"})
}

type SignupInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"e": err.Error()})
		return
	}

	fmt.Println(input.Username)
	fmt.Println(input.Password)

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	_, err := u.CreateUser()
	if err != nil {
		fmt.Println("Failed to create user")
		c.JSON(400, gin.H{"e": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"token": "big daddy"})

}
