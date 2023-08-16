package service

import (
	"fmt"
	"mime/multipart"
	"project-name/config"
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
}

type depositSrv struct {
	awsRepo     repo.AWSRepo
	depositRepo repo.DepositRepo
}

func (c *depositSrv) Add(backImage, frontImage *multipart.FileHeader) (*models.Deposit, *se.ServiceError) {
	var deposit models.Deposit

	err := c.uploadBackImage(&deposit, backImage, config.AppConfig.DEFAULT_USER_ID)
	if err != nil {
		return nil, se.Internal(err, "error when uploading back image")
	}

	err = c.uploadFrontImage(&deposit, frontImage, config.AppConfig.DEFAULT_USER_ID)
	if err != nil {
		return nil, se.Internal(err, "error when uploading back image")
	}

	accountId := config.AppConfig.DEFAULT_ACCOUNT_ID
	depo, errr := c.depositRepo.Add(accountId, &deposit)
	fmt.Println(depo)
	if errr != nil {
		return nil, se.NotFoundOrInternal(err, "deposit not found")
	}

	return depo, nil
}

func (c *depositSrv) Get(depositId string) (*models.Deposit, *se.ServiceError) {
	_, err := uuid.Parse(depositId)
	if err != nil {
		return nil, se.Validating(err)
	}

	deposit, err := c.depositRepo.Get(depositId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "deposit not found")
	}

	return deposit, nil
}

func (c *depositSrv) GetAll() ([]*models.Deposit, *se.ServiceError) {
	deposits, err := c.depositRepo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return deposits, nil
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
