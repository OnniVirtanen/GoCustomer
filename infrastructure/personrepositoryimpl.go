// Package memory is a in-memory implementation of the customer repository
package infrastructure

import (
	"fmt"
	"sync"

	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/repository"
	"github.com/google/uuid"
)

// MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func NewCustomerRepository() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, repository.ErrCustomerNotFound
}

func (mr *MemoryRepository) GetAll() []aggregate.Customer {
	// Create a slice to hold the values
	var customers []aggregate.Customer
	// Iterate over the map and append each value to the slice
	for _, value := range mr.customers {
		customers = append(customers, value)
	}
	return customers
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", repository.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", repository.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Delete removes a customer from the repository
func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	// Make sure Customer is in the repository
	mr.Lock()
	defer mr.Unlock()
	if _, ok := mr.customers[id]; !ok {
		return fmt.Errorf("customer does not exist: %w", repository.ErrCustomerNotFound)
	}

	// Delete the customer
	delete(mr.customers, id)
	return nil
}
