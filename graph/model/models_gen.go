// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type OfficeCreate struct {
	OfficeCode   string  `json:"officeCode"`
	City         string  `json:"city"`
	Phone        string  `json:"phone"`
	AddressLine1 string  `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	State        *string `json:"state"`
	Country      string  `json:"country"`
	PostalCode   string  `json:"postalCode"`
	Territory    string  `json:"territory"`
}

type OfficeUpdate struct {
	OfficeCode   *string `json:"officeCode"`
	City         *string `json:"city"`
	Phone        *string `json:"phone"`
	AddressLine1 *string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	PostalCode   *string `json:"postalCode"`
	Territory    *string `json:"territory"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
