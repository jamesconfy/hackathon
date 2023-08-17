package service

import (
	"fmt"
	"mime/multipart"
	"project-name/config"
	"project-name/internal/forms"
	"project-name/internal/models"
	repo "project-name/internal/repository"
	"project-name/internal/se"
	"strings"

	"github.com/google/uuid"
)

type DepositService interface {
	Add(backImage, frontImage *multipart.FileHeader) (*models.Deposit, *se.ServiceError)
	Get(depositId string) (*models.Deposit, *se.ServiceError)
	GetAll() ([]*models.Deposit, *se.ServiceError)
	Update(depositId string, deposit *forms.Deposit) (*models.Deposit, *se.ServiceError)
}

type depositSrv struct {
	awsRepo     repo.AWSRepo
	depositRepo repo.DepositRepo
}

func (d *depositSrv) Add(backImage, frontImage *multipart.FileHeader) (*models.Deposit, *se.ServiceError) {
	var deposit models.Deposit

	err := d.uploadBackImage(&deposit, backImage, config.AppConfig.DEFAULT_USER_ID)
	if err != nil {
		return nil, se.Internal(err, "error when uploading back image")
	}

	err = d.uploadFrontImage(&deposit, frontImage, config.AppConfig.DEFAULT_USER_ID)
	if err != nil {
		return nil, se.Internal(err, "error when uploading back image")
	}

	depo, errr := d.depositRepo.Add(&deposit)
	if errr != nil {
		return nil, se.Internal(errr)
	}

	return depo, nil
}

func (d *depositSrv) Get(depositId string) (*models.Deposit, *se.ServiceError) {
	_, err := uuid.Parse(depositId)
	if err != nil {
		return nil, se.Validating(err)
	}

	deposit, err := d.depositRepo.Get(depositId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "deposit not found")
	}

	return deposit, nil
}

func (d *depositSrv) GetAll() ([]*models.Deposit, *se.ServiceError) {
	deposits, err := d.depositRepo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return deposits, nil
}

func (d *depositSrv) Update(depositId string, deposit *forms.Deposit) (*models.Deposit, *se.ServiceError) {
	var dep models.Deposit

	_, err := uuid.Parse(depositId)
	if err != nil {
		return nil, se.BadRequest("invalid deposit id")
	}

	dep.Id = depositId
	dep.Status = deposit.Status

	depo, err := d.depositRepo.Update(&dep)
	if err != nil {
		return nil, se.NotFoundOrInternal(err)
	}

	return depo, nil
}

func NewDepositService(awsRepo repo.AWSRepo, depositRepo repo.DepositRepo) DepositService {
	return &depositSrv{awsRepo: awsRepo, depositRepo: depositRepo}
}

// Auxillary Function
func (c *depositSrv) uploadBackImage(deposit *models.Deposit, file *multipart.FileHeader, userId string) error {
	var res models.Image

	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	fileName := fmt.Sprintf("%s/%s.%s", userId, uuid.New().String(), fileType)

	err := c.awsRepo.UploadImage(file, fileName)
	if err != nil {
		return err
	}

	res.Name = fmt.Sprintf("https://check-deposit-bucket.s3.amazonaws.com/%v", fileName)
	res.Size = file.Size
	res.FileType = fileType

	deposit.BackImage = res.Name
	return nil
}

func (c *depositSrv) uploadFrontImage(deposit *models.Deposit, file *multipart.FileHeader, userId string) error {
	var res models.Image

	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[1]

	fileName := fmt.Sprintf("%s/%s.%s", userId, uuid.New().String(), fileType)

	err := c.awsRepo.UploadImage(file, fileName)

	if err != nil {
		return err
	}

	res.Name = fmt.Sprintf("https://check-deposit-bucket.s3.amazonaws.com/%v", fileName)

	res.Size = file.Size
	res.FileType = fileType

	deposit.FrontImage = res.Name
	return nil
}
