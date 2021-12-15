package main

import (
	"duel.io/api/controllers"
	"duel.io/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/users", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)

	//Challenges
	r.GET("/challenges", controllers.FindChallenge)
	r.POST("/challenges", controllers.CreateChallenge)

	//Entries
	r.GET("/entries", controllers.FindEntry)
	r.POST("/entries", controllers.CreateEntry)
	r.POST("/entries/vote", controllers.IncreaseVote)

	// Run the server
	r.Run(":3001")
}
