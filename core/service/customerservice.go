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

func (s *CustomerService) GetAll() ([]aggregate.Customer, error) {
	customers := s.personRepository.GetAll()

	// If customers slice is nil or empty, return an explicit empty slice
	if customers == nil || len(customers) == 0 {
		return []aggregate.Customer{}, nil
	}

	return customers, nil
}

func (s *CustomerService) Get(id uuid.UUID) (aggregate.Customer, error) {
	return s.personRepository.Get(id)
}

func (s *CustomerService) Save(customer aggregate.Customer) error {
	customer.Person.ID = uuid.New()
	s.personRepository.Add(customer)
	return nil
}

func (s *CustomerService) Update(customer aggregate.Customer, id uuid.UUID) error {
	s.personRepository.Update(customer, id)
	return nil
}

func (s *CustomerService) Delete(id uuid.UUID) error {
	s.personRepository.Delete(id)
	return nil
}
