package models

import "time"

type Deposit struct {
	Id          string    `json:"id"`
	BackImage   string    `json:"back_image"`
	FrontImage  string    `json:"front_image"`
	UserId      string    `json:"user_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type Image struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	FileType string `json:"fileType"`
}
