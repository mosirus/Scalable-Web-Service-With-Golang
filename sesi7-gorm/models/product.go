package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("Before insert to table products")

	if len(p.Name) < 4 {
		err = fmt.Errorf("product name to short")
	}

	return
}

func (p *Product) Print() {
	fmt.Println("ID :", p.ID)
	fmt.Println("Brand :", p.Brand)
}
