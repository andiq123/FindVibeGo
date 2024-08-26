package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndRespondBadRequest(model any, context *gin.Context) {
	err := context.ShouldBindBodyWithJSON(&model)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not parse json"})
		return
	}
}
