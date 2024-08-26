package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/ping", s.PingPongHandler)
	r.GET("/searchSongs", s.searchSongs)
	r.GET("/suggestions", s.getSuggestions)

	return r
}
