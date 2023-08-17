package handlers

import (
	"project-name/internal/forms"
	"project-name/internal/response"
	"project-name/internal/se"
	"project-name/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DepositHanlder interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
}

type depositHandler struct {
	depositSrv service.DepositService
}

func (d *depositHandler) Add(c *gin.Context) {
	backImage, err := c.FormFile("back_image")
	if err != nil {
		response.Error(c, *se.BadRequestOrInternal("error when getting front image", err))
		return
	}

	frontImage, err := c.FormFile("front_image")
	if err != nil {
		response.Error(c, *se.BadRequestOrInternal("error when getting front image", err))
		return
	}

	deposit, errr := d.depositSrv.Add(backImage, frontImage)
	if errr != nil {
		response.Error(c, *errr)
		return
	}

	response.Success(c, "Check uploaded successfully", deposit)
}

func (d *depositHandler) Get(c *gin.Context) {
	depositId := c.Param("depositId")
	if depositId == "" {
		response.Error(c, *se.BadRequest("deposit id cannot be null"))
		return
	}

	deposit, errr := d.depositSrv.Get(depositId)
	if errr != nil {
		response.Error(c, *errr)
		return
	}

	response.Success(c, "Deposit gotten successfully", deposit)
}

func (d *depositHandler) GetAll(c *gin.Context) {
	deposit, errr := d.depositSrv.GetAll()
	if errr != nil {
		response.Error(c, *errr)
		return
	}

	response.Success(c, "Deposits gotten successfully", deposit, len(deposit))
}

// Update implements DepositHanlder.
func (d *depositHandler) Update(c *gin.Context) {
	var req forms.Deposit

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.BadRequest("error when getting request"))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	depositId := c.Param("depositId")
	if depositId == "" {
		response.Error(c, *se.BadRequest("deposit id cannot be empty"))
		return
	}

	deposit, err := d.depositSrv.Update(depositId, &req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "status updated successfully", deposit)
}

func NewDepositHandler(depositSrv service.DepositService) DepositHanlder {
	return &depositHandler{depositSrv: depositSrv}
}
