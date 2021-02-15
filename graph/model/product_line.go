package model

type Productline struct {
	ProductLine     string  `json:"productLine" gorm:"column:productLine"`
	TextDescription string  `json:"textDescription" gorm:"column:textDescription"`
	HTMLDescription *string `json:"htmlDescription" gorm:"column:htmlDescription"`
	Image           *string `json:"image" gorm:"column:image"`
}
