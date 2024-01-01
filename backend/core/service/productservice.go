package service

import (
	"fmt"

	"example.com/backend/core/model/entity"
	"example.com/backend/core/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	var productService = ProductService{}
	productService.productRepository = productRepository
	return &productService
}

func (s *ProductService) GetAll() ([]entity.Product, error) {
	products, _ := s.productRepository.GetAll()
	fmt.Println("beginning products at productservice.go GetAll()", products)
	// If product slice is nil or empty, return an explicit empty slice
	if products == nil || len(products) == 0 {
		return []entity.Product{}, nil
	}

	fmt.Println("products at productservice.go GetAll()", products)
	return products, nil
}

func (s *ProductService) Get(id uuid.UUID) (entity.Product, error) {
	return s.productRepository.Get(id)
}

func (s *ProductService) Save(product entity.Product) (entity.Product, error) {
	product.ID = uuid.New()
	err := s.productRepository.Add(product)
	return product, err
}

func (s *ProductService) Update(product entity.Product, id uuid.UUID) error {
	s.productRepository.Update(product, id)
	return nil
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.productRepository.Delete(id)
}
