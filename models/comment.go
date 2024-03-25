package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id" `
	Message string `json:"message" valid:"required~Your message is required"`
	PhotoID uint   `json:"photo_id"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
