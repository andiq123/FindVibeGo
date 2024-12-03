package routes

import (
	"FindVibeGo/cmd/services"
	"FindVibeGo/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var suggestionsService = services.NewSuggestionsService()

func GetSuggestionsHandler(context *gin.Context) {
	searchQuery := context.Query("q")
	searchQuery, err := utils.CleanString(searchQuery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	suggestions, err := suggestionsService.GetSuggestions(searchQuery)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"result": suggestions})
}
