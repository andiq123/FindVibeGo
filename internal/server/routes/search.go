package routes

import (
	"FindVibeGo/cmd/services"
	"FindVibeGo/cmd/utils"

	"github.com/gin-gonic/gin"
)

var muzkanService = services.NewMuzkanScrapperService()

func SearchSongsHandler(context *gin.Context) {
	searchQuery := context.Query("q")
	searchQuery, err := utils.CleanString(searchQuery)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	songs, err := muzkanService.GetSongs(searchQuery)
	if err != nil {
		context.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if len(songs) == 0 {
		context.JSON(404, gin.H{"message": "No songs found by this query"})
		return
	}

	context.JSON(200, gin.H{"result": songs})
}
