package model

type Customer struct {
	CustomerNumber   int      `json:"customerNumber" gorm:"column:customerNumber"`
	CustomerName     string   `json:"customerName" gorm:"column:customerName"`
	ContactLastName  string   `json:"contactLastName" gorm:"column:contactLastName"`
	ContactFirstName string   `json:"contactFirstName" gorm:"column:contactFirstName"`
	Phone            string   `json:"phone" gorm:"column:phone"`
	AddressLine1     string   `json:"addressLine1" gorm:"column:addressLine1"`
	AddressLine2     *string  `json:"addressLine2" gorm:"column:addressLine2"`
	City             string   `json:"city" gorm:"column:city"`
	State            *string  `json:"state" gorm:"column:state"`
	PostalCode       *string  `json:"postalCode" gorm:"column:postalCode"`
	Country          string   `json:"country" gorm:"column:country"`
	EmployeeID       *int     `json:"employee" gorm:"column:salesRepEmployeeNumber"`
	CreditLimit      *float64 `json:"creditLimit" gorm:"column:creditLimit"`
}
