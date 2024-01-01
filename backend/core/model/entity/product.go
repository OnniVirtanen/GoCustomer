package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uuid.UUID       `json:"ID"`
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Quantity    uint            `json:"quantity" binding:"required"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Image       string          `json:"image" binding:"required"`
	Category    string          `json:"category" binding:"required"`
	CreatedAt   *time.Time      `json:"createdAt"`
	UpdatedAt   *time.Time      `json:"updatedAt"`
	Discount    decimal.Decimal `json:"discount" binding:"required"`
}
