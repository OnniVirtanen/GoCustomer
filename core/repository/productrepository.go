// Package Customer holds all the domain logic for the customer domain.
package repository

import (
	"errors"

	"example.com/backend/core/model/entity"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("the product was not found in the repository")
	ErrFailedToAddProduct = errors.New("failed to add the product to the repository")
	ErrUpdateProduct      = errors.New("failed to update the product in the repository")
)

// CustomerRepository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type ProductRepository interface {
	Get(uuid.UUID) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Add(entity.Product) error
	Update(entity.Product, uuid.UUID) error
	Delete(uuid.UUID) error
}
