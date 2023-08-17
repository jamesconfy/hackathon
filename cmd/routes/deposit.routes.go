package routes

import (
	"project-name/cmd/handlers"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
)

func DepositRoute(router *gin.RouterGroup, depositService service.DepositService) {
	handler := handlers.NewDepositHandler(depositService)
	deposit := router.Group("/deposits/cheque")
	{
		deposit.POST("", handler.Add)
		deposit.GET("", handler.GetAll)
		deposit.GET("/:depositId", handler.Get)
		deposit.PATCH("/:depositId", handler.Update)
	}

}
