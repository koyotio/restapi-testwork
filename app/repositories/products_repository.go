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
	Id         int    `db:"id"`
	CategoryId int    `db:"category_id"`
	Articul    string `db:"articul"`
	Name       string `db:"name"`
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
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if productInput.CategoryId != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"category_id"=$%d`, numberOfArguments))
		args = append(args, *productInput.CategoryId)
		product.CategoryId = *productInput.CategoryId
	}

	if productInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"name"=$%d`, numberOfArguments))
		args = append(args, *productInput.Name)
		product.Name = *productInput.Name
	}

	if productInput.Articul != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"articul"=$%d`, numberOfArguments))
		args = append(args, *productInput.Articul)
		product.Articul = *productInput.Articul
	}

	insertSet := strings.Join(values, `, `)

	err := p.db.QueryRow(
		fmt.Sprintf(`INSERT INTO "products" SET %s RETURNING "id"`, insertSet),
		args...,
	).Scan(&product.Id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p DBProducts) Update(id int, productInput input.ProductInput) (*Product, error) {
	product := Product{
		Id: id,
	}
	values := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if productInput.CategoryId != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"category_id"=$%d`, numberOfArguments))
		args = append(args, *productInput.CategoryId)
		product.CategoryId = *productInput.CategoryId
	}

	if productInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"name"=$%d`, numberOfArguments))
		args = append(args, *productInput.Name)
		product.Name = *productInput.Name
	}

	if productInput.Articul != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"articul"=$%d`, numberOfArguments))
		args = append(args, *productInput.Articul)
		product.Articul = *productInput.Articul
	}

	updateSet := strings.Join(values, `, `)
	args = append(args, id)
	numberOfArguments++

	result, err := p.db.Exec(
		fmt.Sprintf(`UPDATE "products" SET %s WHERE "id"=$%d`, updateSet, numberOfArguments),
		args...,
	)
	if result == nil {
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
	product := Product{Id: id}
	row := p.db.QueryRow(
		`SELECT "category_id", "name", "articul" FROM "products" WHERE "id"=$1`,
		id,
	).Scan(
		&product.CategoryId,
		&product.Articul,
		&product.Name,
	)
	if row == nil {
		return nil, errors.New(`product not found`)
	}
	if len(row.Error()) != 0 {
		return nil, errors.New(row.Error())
	}
	return &product, nil
}
