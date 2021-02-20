package services

import (
	"rest-api/input"
	"rest-api/repositories"
)

type ProductServiceConfig struct {
	RepositoryManager *repositories.RepositoryManager
}

type ProductService interface {
	CreateProduct(productInput input.ProductInput) (*repositories.Product, error)
	UpdateProduct(id int, productInput input.ProductInput) (*repositories.Product, error)
	DeleteProduct(id int) (bool, error)
	GetAllProducts() ([]repositories.Product, error)
	GetProductById(id int) (*repositories.Product, error)
}

func NewProductService(repositoryManager *repositories.RepositoryManager) *ProductServiceConfig {
	return &ProductServiceConfig{RepositoryManager: repositoryManager}
}

func (productService *ProductServiceConfig) CreateProduct(productInput input.ProductInput) (*repositories.Product, error) {
	return productService.RepositoryManager.ProductsRepository.Create(productInput)
}

func (productService *ProductServiceConfig) UpdateProduct(id int, productInput input.ProductInput) (*repositories.Product, error) {
	return productService.RepositoryManager.ProductsRepository.Update(id, productInput)
}

func (productService *ProductServiceConfig) DeleteProduct(id int) (bool, error) {
	return productService.RepositoryManager.ProductsRepository.Delete(id)
}

func (productService *ProductServiceConfig) GetAllProducts() ([]repositories.Product, error) {
	return productService.RepositoryManager.ProductsRepository.GetAll()
}

func (productService *ProductServiceConfig) GetProductById(id int) (*repositories.Product, error) {
	return productService.RepositoryManager.ProductsRepository.GetById(id)
}
