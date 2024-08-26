package routes

import (
	"FindVibeGo/cmd/services"
	"FindVibeGo/cmd/utils"
	"github.com/gin-gonic/gin"
)

func SearchSongsHandler(context *gin.Context) {
	searchQuery := context.Query("q")
	searchQuery, err := utils.CleanString(searchQuery)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	muzkanService := services.NewMuzkanScrapperService()
	songs, err := muzkanService.GetSongs(searchQuery)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, songs)
}
