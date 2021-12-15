package controllers

import (
	"net/http"

	"duel.io/api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

// GET /users
// Find all books
func FindUser(c *gin.Context) {
	var user []models.User
	models.DB.Find(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
