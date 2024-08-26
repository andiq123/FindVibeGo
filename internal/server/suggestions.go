package server

import (
	"FindVibeGo/cmd/services"
	"FindVibeGo/cmd/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getSuggestions(context *gin.Context) {
	searchQuery := context.Query("q")
	searchQuery, err := utils.CleanString(searchQuery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	suggestionsService := services.NewSuggestionsService()
	suggestions, err := suggestionsService.GetSuggestions(searchQuery)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(200, suggestions)
}
