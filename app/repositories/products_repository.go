package repositories

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-api/input"
	"strings"
)

type DBProducts struct {
	db *sqlx.DB
}

func NewDBProducts(db *sqlx.DB) *DBProducts {
	return &DBProducts{db: db}
}

type Product struct {
	Id         int    `db:"id" json:"id"`
	CategoryId int    `db:"category_id" json:"category_id"`
	Articul    string `db:"articul" json:"articul"`
	Name       string `db:"name" json:"name"`
}

type ProductsRepository interface {
	Create(productInput input.ProductInput) (*Product, error)
	Update(id int, productInput input.ProductInput) (*Product, error)
	Delete(id int) (bool, error)
	GetAll() ([]Product, error)
	GetById(id int) (*Product, error)
}

func (p DBProducts) Create(productInput input.ProductInput) (*Product, error) {
	product := Product{}
	values := make([]string, 0)
	columns := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if productInput.CategoryId != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`$%d`, numberOfArguments))
		columns = append(columns, `"category_id"`)
		args = append(args, *productInput.CategoryId)
	}

	if productInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`$%d`, numberOfArguments))
		columns = append(columns, `"name"`)
		args = append(args, *productInput.Name)
	}

	if productInput.Articul != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`$%d`, numberOfArguments))
		columns = append(columns, `"articul"`)
		args = append(args, *productInput.Articul)
	}

	err := p.db.QueryRow(
		fmt.Sprintf(
			`INSERT INTO "products"  (%s) VALUES(%s) RETURNING "id", "name", "articul", "category_id"`,
			strings.Join(columns, `, `),
			strings.Join(values, `, `),
		),
		args...,
	).Scan(&product.Id, &product.Name, &product.Articul, &product.CategoryId)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p DBProducts) Update(id int, productInput input.ProductInput) (*Product, error) {
	product := Product{}
	values := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if productInput.CategoryId != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"category_id"=$%d`, numberOfArguments))
		args = append(args, *productInput.CategoryId)
	}

	if productInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"name"=$%d`, numberOfArguments))
		args = append(args, *productInput.Name)
	}

	if productInput.Articul != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"articul"=$%d`, numberOfArguments))
		args = append(args, *productInput.Articul)
	}

	updateSet := strings.Join(values, `, `)
	args = append(args, id)
	numberOfArguments++

	row := p.db.QueryRow(
		fmt.Sprintf(
			`UPDATE "products" SET %s WHERE "id"=$%d RETURNING "id", "name", "articul", "category_id"`,
			updateSet,
			numberOfArguments,
		),
		args...,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(&product.Id, &product.Name, &product.Articul, &product.CategoryId)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (p DBProducts) Delete(id int) (bool, error) {
	result := p.db.MustExec(`DELETE FROM "products" WHERE "id"=$1`, id)
	affectedRows, err := result.RowsAffected()
	return affectedRows == 1, err
}

func (p DBProducts) GetAll() ([]Product, error) {
	var products []Product
	err := p.db.Select(&products, `SELECT "id", "category_id", "name", "articul" FROM "products"`)
	return products, err
}

func (p DBProducts) GetById(id int) (*Product, error) {
	product := Product{}
	err := p.db.QueryRow(
		`SELECT "id", "category_id", "name", "articul" FROM "products" WHERE "id"=$1`,
		id,
	).Scan(
		&product.Id,
		&product.CategoryId,
		&product.Articul,
		&product.Name,
	)

	if err != nil {
		return nil, errors.New(`product not found`)
	}
	return &product, nil
}
