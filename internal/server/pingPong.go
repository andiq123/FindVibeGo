package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) PingPongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
