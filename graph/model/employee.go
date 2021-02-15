package model

type Employee struct {
	EmployeeNumber int    `json:"employeeNumber" gorm:"column:employeeNumber"`
	LastName       string `json:"lastName" gorm:"column:lastName"`
	FirstName      string `json:"firstName" gorm:"column:firstName"`
	Extension      string `json:"extension" gorm:"column:extension"`
	Email          string `json:"email" gorm:"column:email"`
	OfficeID       string `json:"office" gorm:"column:officeCode"`
	EmployeeID     *int   `json:"employee" gorm:"column:reportsTo"`
	JobTitle       string `json:"jobTitle" gorm:"column:jobTitle"`
}
