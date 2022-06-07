package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId  uint   `form:"user_id" json:"user_id" form:"user_id"`
	PhotoId uint   `form:"photo_id" json:"photo_id" form:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	User    User   `valid:"-"`
	Photo   Photo  `valid:"-"`
}

type CommentResponse struct {
	GormModel
	UserId  uint   `form:"user_id" json:"user_id"`
	PhotoId uint   `form:"photo_id" json:"photo_id"`
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~message is required"`
	User    UserProfileId
	Photo   ResponsePhoto
}

type CommentResponseUpdate struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
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
