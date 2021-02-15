package model

type Orderdetail struct {
	OrderID         int     `json:"order" gorm:"column:orderNumber"`
	ProductID       string  `json:"product" gorm:"column:productCode"`
	QuantityOrdered int     `json:"quantityOrdered" gorm:"column:quantityOrdered"`
	PriceEach       float64 `json:"priceEach" gorm:"column:priceEach"`
	OrderLineNumber int     `json:"orderLineNumber" gorm:"column:orderLineNumber"`
}
