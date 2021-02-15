package model

type Office struct {
	OfficeCode   string  `json:"officeCode" gorm:"column:officeCode" gorm:"primaryKey"`
	City         string  `json:"city" gorm:"column:city"`
	Phone        string  `json:"phone" gorm:"column:phone"`
	AddressLine1 string  `json:"addressLine1" gorm:"column:addressLine1"`
	AddressLine2 *string `json:"addressLine2" gorm:"column:addressLine2"`
	State        *string `json:"state" gorm:"column:state"`
	Country      string  `json:"country" gorm:"column:country"`
	PostalCode   string  `json:"postalCode" gorm:"column:postalCode"`
	Territory    string  `json:"territory" gorm:"column:territory"`
}
