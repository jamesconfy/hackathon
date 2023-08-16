package handlers

import (
	"project-name/internal/response"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
)

type HomeHandler interface {
	Home(c *gin.Context)
}

type homeHandler struct {
	homeSrv service.HomeService
}

// Home godoc
// @Summary	Get home route
// @Description	Get home details
// @Tags	Home
// @Accept	json
// @Produce	json
// @Success	200  {object}	response.SuccessMessage
// @Failure	400  {object}  se.ServiceError
// @Failure	404  {object}  se.ServiceError
// @Failure	500  {object}  se.ServiceError
// @Router	/ [get]
func (h *homeHandler) Home(c *gin.Context) {
	message, err := h.homeSrv.CreateHome()
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, message, nil)
}

func NewHomeHandler(homeSrv service.HomeService) HomeHandler {
	return &homeHandler{homeSrv: homeSrv}
}
