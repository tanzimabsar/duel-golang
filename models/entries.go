package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Link        string `json:"link" binding:"required"`
	ChallengeID int    `json:"challenge_id"`
	Votes       int    `json:"votes" gorm:"default:0"`
}
