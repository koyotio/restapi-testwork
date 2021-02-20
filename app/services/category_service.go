package services

import (
	"rest-api/input"
	"rest-api/repositories"
)

type CategoryServiceConfig struct {
	RepositoryManager *repositories.RepositoryManager
}

type CategoryService interface {
	CreateCategory(categoryInput *input.CategoryInput) (*repositories.Category, error)
	UpdateCategory(id int, categoryInput *input.CategoryInput) (*repositories.Category, error)
	DeleteCategory(id int) (bool, error)
	GetAllCategories() ([]repositories.Category, error)
	GetCategoryById(id int) (*repositories.Category, error)
}

func NewCategoryService(repositoryManager *repositories.RepositoryManager) *CategoryServiceConfig {
	return &CategoryServiceConfig{
		RepositoryManager: repositoryManager,
	}
}

func (categoryService *CategoryServiceConfig) CreateCategory(categoryInput *input.CategoryInput) (*repositories.Category, error) {
	return categoryService.RepositoryManager.CategoriesRepository.Create(categoryInput)
}

func (categoryService *CategoryServiceConfig) UpdateCategory(id int, categoryInput *input.CategoryInput) (*repositories.Category, error) {
	return categoryService.RepositoryManager.CategoriesRepository.Update(id, categoryInput)
}

func (categoryService *CategoryServiceConfig) DeleteCategory(id int) (bool, error) {
	return categoryService.RepositoryManager.CategoriesRepository.Delete(id)
}

func (categoryService *CategoryServiceConfig) GetAllCategories() ([]repositories.Category, error) {
	return categoryService.RepositoryManager.CategoriesRepository.GetAll()
}

func (categoryService *CategoryServiceConfig) GetCategoryById(id int) (*repositories.Category, error) {
	return categoryService.RepositoryManager.CategoriesRepository.GetById(id)
}
