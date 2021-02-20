package repositories

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-api/input"
	"strings"
)

type DBCategories struct {
	db *sqlx.DB
}

type Category struct {
	Id   int          `db:"id" json:"id"`
	Name string       `db:"name" json:"name"`
	Tags CategoryTags `db:"tags" json:"tags"`
}

type CategoryTags []string

type CategoriesRepository interface {
	Create(categoryInput *input.CategoryInput) (*Category, error)
	Update(id int, categoryInput *input.CategoryInput) (*Category, error)
	Delete(id int) (bool, error)
	GetAll() ([]Category, error)
	GetById(categoryId int) (*Category, error)
}

func NewDBCategories(db *sqlx.DB) *DBCategories {
	return &DBCategories{db: db}
}

func (c *DBCategories) Create(categoryInput *input.CategoryInput) (*Category, error) {
	category := Category{}
	values := make([]string, 0)
	columns := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if categoryInput.Tags != nil {
		encodedTags, err := json.Marshal(*categoryInput.Tags)
		if err != nil {
			return nil, errors.New("json encode error when saving to db")
		}
		numberOfArguments++
		values = append(values, fmt.Sprintf(`$%d`, numberOfArguments))
		columns = append(columns, `"tags"`)
		args = append(args, string(encodedTags))
	}

	if categoryInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`$%d`, numberOfArguments))
		columns = append(columns, `"name"`)
		args = append(args, *categoryInput.Name)
	}

	err := c.db.QueryRow(
		fmt.Sprintf(
			`INSERT INTO "categories" (%s) VALUES(%s) RETURNING "id", "name", "tags"`,
			strings.Join(columns, `, `),
			strings.Join(values, `, `),
		),
		args...,
	).Scan(&category.Id, &category.Name, &category.Tags)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func (c *DBCategories) Update(id int, categoryInput *input.CategoryInput) (*Category, error) {
	category := Category{Id: id}
	values := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if categoryInput.Tags != nil {
		encodedTags, err := json.Marshal(*categoryInput.Tags)
		if err != nil {
			return nil, errors.New("json encode error when saving to db")
		}
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"tags"=$%d`, numberOfArguments))
		args = append(args, encodedTags)
		category.Tags = *categoryInput.Tags
	}

	if categoryInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"name"=$%d`, numberOfArguments))
		args = append(args, *categoryInput.Name)
		category.Name = *categoryInput.Name
	}

	updateSet := strings.Join(values, `, `)
	args = append(args, id)
	numberOfArguments++

	row := c.db.QueryRow(
		fmt.Sprintf(
			`UPDATE "categories" SET %s WHERE "id"=$%d RETURNING "id", "name", "tags"`,
			updateSet,
			numberOfArguments,
		),
		args...,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(&category.Id, &category.Name, &category.Tags)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func (c *DBCategories) Delete(id int) (bool, error) {
	result := c.db.MustExec(`DELETE FROM "categories" WHERE "id"=$1`, id)
	affectedRows, err := result.RowsAffected()
	return affectedRows == 1, err
}

func (c *DBCategories) GetAll() ([]Category, error) {
	var categories []Category
	err := c.db.Select(&categories, `SELECT "id", "name", "tags" FROM "categories"`)
	return categories, err
}

func (c *DBCategories) GetById(categoryId int) (*Category, error) {
	category := Category{}
	err := c.db.QueryRow(
		`SELECT "id", "name", "tags" FROM "categories" WHERE "id"=$1`,
		categoryId,
	).Scan(
		&category.Id, &category.Name, &category.Tags,
	)

	if err != nil {
		return nil, errors.New(`category not found`)
	}
	return &category, nil
}

func (categoryTags *CategoryTags) Value() (driver.Value, error) {
	jsonCategoryTags, err := json.Marshal(categoryTags)
	return jsonCategoryTags, err
}

func (categoryTags *CategoryTags) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		source := src.([]byte)
		err := json.Unmarshal(source, &categoryTags)
		if err != nil {
			return err
		}
		return nil
	case []string:
		return nil
	default:
		return errors.New("unsupported variable type")
	}

}
