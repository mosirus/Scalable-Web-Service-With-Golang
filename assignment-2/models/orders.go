package models

import "time"

type Order struct {
	ID           int       `gorm:"primary_key" json:"orderId"`
	CustomerName string    `json:"customerName" gorm:"not null;type:varchar(191)"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
}

type CreateOrder struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
