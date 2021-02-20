package services

import "rest-api/repositories"

type ServiceManager struct {
	CategoryService
	ProductService
}

func NewServiceManager(repositoryManager *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CategoryService: NewCategoryService(repositoryManager),
		ProductService:  NewProductService(repositoryManager),
	}
}
