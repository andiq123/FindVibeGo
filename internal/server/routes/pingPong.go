package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingPongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
