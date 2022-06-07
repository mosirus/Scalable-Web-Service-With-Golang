package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url is required"`
	UserId         uint
	User           User `valid:"-"`
}

type ResponseSocialMedia struct {
	GormModel
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint
	User           UserProfileId
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
