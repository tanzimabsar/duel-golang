package controllers

import (
	"net/http"

	"duel.io/api/models"
	"github.com/gin-gonic/gin"
)

type CreateEntriesInput struct {
	Name        string `json:"name" binding:"required"`
	Link        string `json:"link" binding:"required"`
	ChallengeID int    `json:"challenge_id" binding:"required" gorm:"references":ID`
}

type UpdateVotesInput struct {
	ID int `json:"entry_id" binding:"required"`
}

// GET /users
// Find all books
func FindEntry(c *gin.Context) {
	var entry []models.Entry
	models.DB.Find(&entry)

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

func CreateEntry(c *gin.Context) {
	// Validate input
	var input CreateEntriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	entry := models.Entry{
		Name:        input.Name,
		Link:        input.Link,
		ChallengeID: input.ChallengeID,
	}

	models.DB.Create(&entry)

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

//Update Method to increase vote by 1

func IncreaseVote(c *gin.Context) {
	var input UpdateVotesInput
	// Create book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var entry models.Entry
	models.DB.Find(&entry, input.ID)
	entry.Votes = entry.Votes + 1
	models.DB.Save(&entry)
	c.JSON(http.StatusOK, gin.H{"data": entry})

}
