package models

import (
	"errors"
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null, uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email       string        `gorm:"not null , uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~password is required, minstringlength(6)~minimum length password is 6"`
	Age         int           `gorm:"not null" json:"age" form:"age" valid:"required~age is required"`
	Comment     []Comment     `gorm:"foreignKey:UserId" valid:"-"`
	Photo       []Photo       `gorm:"foreignKey:UserId" valid:"-"`
	SocialMedia []SocialMedia `gorm:"foreignKey:UserId" valid:"-"`
}

type UserProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserProfileId struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	if u.Age < 8 {
		err = errors.New("minimum Age is 8")
		return err
	}
	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
