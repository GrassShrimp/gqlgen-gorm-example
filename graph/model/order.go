package model

import "time"

type Order struct {
	OrderNumber  int        `json:"orderNumber" gorm:"column:orderNumber"`
	OrderDate    time.Time  `json:"orderDate" gorm:"column:orderDate"`
	RequiredDate time.Time  `json:"requiredDate" gorm:"column:requiredDate"`
	ShippedDate  *time.Time `json:"shippedDate" gorm:"column:shippedDate"`
	Status       string     `json:"status" gorm:"column:status"`
	Comments     *string    `json:"comments" gorm:"column:comments"`
	CustomerID   int        `json:"customer" gorm:"column:customerNumber"`
}
