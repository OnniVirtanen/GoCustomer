// Package Customer holds all the domain logic for the customer domain.
package repository

import (
	"errors"

	"example.com/backend/core/model/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	GetAll() []aggregate.Customer
	Add(aggregate.Customer) error
	Update(aggregate.Customer, uuid.UUID) error
	Delete(uuid.UUID) error
}
