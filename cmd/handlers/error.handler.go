package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error404() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"description": "Page not found.",
			"status":      http.StatusNotFound,
		})
	}
}
