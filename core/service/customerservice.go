package service

import (
	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/repository"
	"github.com/google/uuid"
)

type CustomerService struct {
	personRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) *CustomerService {
	var customerService = CustomerService{}
	customerService.personRepository = customerRepository
	return &customerService
}

func (cs *CustomerService) GetAll() ([]aggregate.Customer, error) {
	customers := cs.personRepository.GetAll()

	// If customers slice is nil or empty, return an explicit empty slice
	if customers == nil || len(customers) == 0 {
		return []aggregate.Customer{}, nil
	}

	return customers, nil
}

func (cs *CustomerService) Get(id uuid.UUID) (aggregate.Customer, error) {
	return cs.personRepository.Get(id)
}

func (cs *CustomerService) Save(customer aggregate.Customer) error {
	customer.Person.ID = uuid.New()
	cs.personRepository.Add(customer)
	return nil
}

func (cs *CustomerService) Update(customer aggregate.Customer, id uuid.UUID) error {
	cs.personRepository.Update(customer, id)
	return nil
}

func (cs *CustomerService) Delete(id uuid.UUID) error {
	cs.personRepository.Delete(id)
	return nil
}
