package model

type Product struct {
	ProductCode        string  `json:"productCode" gorm:"column:productCode"`
	ProductName        string  `json:"productName" gorm:"column:productName"`
	ProductLineID      string  `json:"productLine" gorm:"column:productLine"`
	ProductScale       string  `json:"productScale" gorm:"column:productScale"`
	ProductVendor      string  `json:"productVendor" gorm:"column:productVendor"`
	ProductDescription string  `json:"productDescription" gorm:"column:productDescription"`
	QuantityInStock    int     `json:"quantityInStock" gorm:"column:quantityInStock"`
	BuyPrice           float64 `json:"buyPrice" gorm:"column:buyPrice"`
	Msrp               float64 `json:"MSRP" gorm:"column:MSRP"`
}
