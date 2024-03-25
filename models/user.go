package models

import (
	"final_project_golang/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username        string        `gorm:"not null; uniqueIndex" valid:"required~Your name is required" json:"username"`
	Email           string        `gorm:"not null; uniqueIndex" valid:"required~Your email is required, email~Invalid email format" json:"email"`
	Password        string        `json:"password" gorm:"not null" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age             int           `json:"age" gorm:"not null" valid:"required~Your age is required, range(8|120)~Age must be 8 or more"`
	ProfileImageURL string        `json:"profile_image_url,omitempty" form:"profile_image_url"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HassPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fields := map[string]interface{}{
		"Email":    u.Email,
		"Username": u.Username,
	}

	if _, err := govalidator.ValidateMap(fields, map[string]interface{}{"Email": "required,email", "Username": "required"}); err != nil {
		return err
	}

	return nil
}
