package repo

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSRepo interface {
	UploadImage(file *multipart.FileHeader, filename string) error
}

type awsRepo struct {
	s3session *s3.S3
}

func (a *awsRepo) UploadImage(file *multipart.FileHeader, filename string) error {
	image, err := file.Open()
	if err != nil {
		return err
	}

	_, err = a.s3session.PutObject(&s3.PutObjectInput{
		Body:        image,
		Bucket:      aws.String("check-deposit-bucket"),
		Key:         aws.String(filename),
		ContentType: aws.String("image/jpeg"),
	})

	if err != nil {
		return err
	}

	return nil
}

func NewAWSRepo(s3session *s3.S3) AWSRepo {
	return &awsRepo{s3session: s3session}
}

// Auxillary Function
// func (a *awsRepo) resizeImage(image multipart.File) (image.Image, error) {
// 	var imgs []image.Image
// 	var base64Strings []string

// 	fl, err := image.Open()
// 	if err != nil {
// 		return err
// 	}

// 	defer fl.Close()

// 	flRead, err := ioutil.ReadAll(fl)
// 	if err != nil {
// 		log.Println("Couldnot open file")
// 		return err
// 	}

// 	m, err := jpeg.Decode(bytes.NewReader(flRead))
// 	if err != nil {
// 		log.Println("Couldnot decode image")
// 		return base64Strings, err
// 	}
// 	resizedImage := utils.ResizeImg(m)
// 	imgs = append(imgs, resizedImage)

// 	return nil
// }
