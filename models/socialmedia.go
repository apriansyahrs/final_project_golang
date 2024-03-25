package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" valid:"required~Your name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Your social media url is required"`
	UserID         uint   `json:"user_id"`
	User           *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
