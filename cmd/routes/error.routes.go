package routes

import (
	"project-name/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func ErrorRoute(router *gin.Engine) {
	router.NoRoute(handlers.Error404())
}
