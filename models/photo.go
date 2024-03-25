package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" valid:"required~Your title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~Your photo url is required"`
	UserID   uint   `json:"user_id"`
	User     *User
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
