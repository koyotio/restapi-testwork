package repositories

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryManager struct {
	CategoriesRepository
	ProductsRepository
}

func NewRepositoryDBManager(db *sqlx.DB) *RepositoryManager {
	return &RepositoryManager{
		CategoriesRepository: NewDBCategories(db),
		ProductsRepository:   NewDBProducts(db),
	}
}
