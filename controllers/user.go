package controllers

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
