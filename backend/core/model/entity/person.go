package entity

import (
	"github.com/google/uuid"
)

type Person struct {
	ID          uuid.UUID `json:"ID"`
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone" binding:"required,e164"`
	CountryCode string    `json:"countryCode" binding:"required,iso3166_1_alpha2"`
	Street      string    `json:"street" binding:"required"`
	PostalCode  string    `json:"postalCode" binding:"required"`
	City        string    `json:"city" binding:"required"`
	Age         uint      `json:"age" binding:"required"`
}
