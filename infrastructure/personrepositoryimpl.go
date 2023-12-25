// Package memory is a in-memory implementation of the customer repository
package infrastructure

import (
	"fmt"
	"log"
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
	mr.Lock()
	defer mr.Unlock()

	// Create a slice to hold the values
	var customers []aggregate.Customer
	log.Println("Customers:", customers)
	// Iterate over the map and append each value to the slice
	for _, customer := range mr.customers {
		customers = append(customers, customer)
	}

	log.Println("Customers:", customers)
	return customers
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	log.Println("beginning of Add:", c)
	mr.Lock()
	defer mr.Unlock()

	// Safety check if customers map is not created
	if mr.customers == nil {
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
	}

	// Check if Customer's person field is nil
	if c.Person == nil {
		return fmt.Errorf("cannot add customer: person data is missing")
	}

	// Make sure Customer isn't already in the repository
	id := c.Person.ID
	if _, ok := mr.customers[id]; ok {
		return fmt.Errorf("customer already exists: %w", repository.ErrFailedToAddCustomer)
	}
	log.Println("Customer:", c)
	mr.customers[id] = c
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c aggregate.Customer, id uuid.UUID) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[id]; !ok {
		return fmt.Errorf("customer does not exist: %w", repository.ErrUpdateCustomer)
	}
	c.Person.ID = id
	mr.Lock()
	mr.customers[id] = c
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
