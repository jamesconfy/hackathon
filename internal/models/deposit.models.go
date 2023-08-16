package models

import (
	"fmt"
	"time"
)

type Deposit struct {
	Id            string    `json:"id"`
	BackImage     string    `json:"back_image"`
	FrontImage    string    `json:"front_image"`
	AccountNumber string    `json:"account_number"`
	UserId        string    `json:"-"`
	Status        string    `json:"status"`
	User          *User     `json:"user"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
}

type Image struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	FileType string `json:"fileType"`
}

type User struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"-"`
	LastName    string    `json:"-"`
	FullName    string    `json:"name"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

func (u *User) GetFullName() {
	u.FullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
