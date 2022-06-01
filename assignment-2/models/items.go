package models

type Item struct {
	ID          int    `gorm:"primary_key" json:"LineItemId"`
	ItemCode    string `json:"itemCode" gorm:"not null;type:varchar(191)"`
	Description string `json:"description" gorm:"not null;type:varchar(191)"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}
