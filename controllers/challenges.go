package controllers

import (
	"net/http"

	"duel.io/api/models"
	"github.com/gin-gonic/gin"
)

type CreateChallengeInput struct {
	Name string `json:"name" binding:"required"`
}

// GET /users
// Find all books
func FindChallenge(c *gin.Context) {
	var challenge []models.Challenge
	models.DB.Find(&challenge)

	c.JSON(http.StatusOK, gin.H{"data": challenge})
}

func CreateChallenge(c *gin.Context) {
	// Validate input
	var input CreateChallengeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	challenge := models.Challenge{
		Name: input.Name,
	}

	models.DB.Create(&challenge)

	c.JSON(http.StatusOK, gin.H{"data": challenge})
}
