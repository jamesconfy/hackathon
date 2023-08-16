package routes

import (
	"project-name/cmd/handlers"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.RouterGroup, homeSrv service.HomeService) {
	handler := handlers.NewHomeHandler(homeSrv)

	router.Group("").GET("", handler.Home)
}
