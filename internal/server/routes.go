package server

import (
	"FindVibeGo/internal/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/ping", routes.PingPongHandler)
	r.GET("/searchSongs", routes.SearchSongsHandler)
	r.GET("/suggestions", routes.GetSuggestionsHandler)

	return r
}
