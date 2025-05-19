package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Title          string `json:"title"`
	PhoneNumber    string `json:"phone_number"`
	About          string `json:"about" gorm:"type:text"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Document       string `json:"document"`
}

func (User) TableName() string {
	return "users"
}
