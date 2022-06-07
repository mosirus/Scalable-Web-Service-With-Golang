package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption  string    `gorm:"not null" json:"caption" form:"caption" valid:"required~caption is required"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	Comment  []Comment `gorm:"foreignKey:PhotoId" json:"comments" valid:"-"`
	UserId   uint
	User     User `valid:"-"`
}

type APIPhoto struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~caption is required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	UserId   uint
	User     UserProfile
}

type ResponsePhoto struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~caption is required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	UserId   uint   `json:"user_id"`
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
