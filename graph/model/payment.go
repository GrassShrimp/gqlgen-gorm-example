package model

import "time"

type Payment struct {
	Customer    *Customer `json:"customer" gorm:"references:CustomerNumber"`
	CustomerID  int       `gorm:"column:customerNumber"`
	CheckNumber string    `json:"checkNumber" gorm:"column:checkNumber"`
	PaymentDate time.Time `json:"paymentDate" gorm:"column:paymentDate"`
	Amount      float64   `json:"amount" gorm:"column:amount"`
}
