package service

import (
	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/repository"
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
	return s.personRepository.GetAll(), nil
}

/*
func (s *FoodService) SaveFood(foodVM *entity.FoodViewModel) (*entity.FoodViewModel, error) {

	var food = entity.Food{
		UserID:      foodVM.UserID,
		Title:       foodVM.Title,
		Description: foodVM.Description,
		FoodImage:   foodVM.FoodImage,
	}

	result, err := s.foodRepo.SaveFood(&food)

	if err != nil {
		return nil, err
	}

	if result != nil {
		foodVM = &entity.FoodViewModel{
			ID:          result.ID,
			UserID:      result.UserID,
			Title:       result.Title,
			Description: result.Description,
			FoodImage:   result.FoodImage,
		}
	}

	return foodVM, nil
}
*/
