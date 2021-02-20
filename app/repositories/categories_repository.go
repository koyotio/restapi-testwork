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
	Id   int          `db:"id"`
	Name string       `db:"name"`
	Tags CategoryTags `db:"tags"`
}

type CategoryTags []string

type CategoriesRepository interface {
	Create(categoryInput input.CategoryInput) (*Category, error)
	Update(id int, categoryInput input.CategoryInput) (*Category, error)
	Delete(id int) (bool, error)
	GetAll() ([]Category, error)
	GetById(categoryId int) (*Category, error)
}

func NewDBCategories(db *sqlx.DB) *DBCategories {
	return &DBCategories{db: db}
}

func (c *DBCategories) Create(categoryInput input.CategoryInput) (*Category, error) {
	category := Category{}
	values := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if categoryInput.Tags != nil {
		category.Tags = CategoryTags{}
		err := category.Tags.Scan(*categoryInput.Tags)
		if err != nil {
			return nil, err
		}
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"tags"=$%d`, numberOfArguments))
		args = append(args, category.Tags)
	}

	if categoryInput.Name != nil {
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"name"=$%d`, numberOfArguments))
		args = append(args, *categoryInput.Name)
		category.Name = *categoryInput.Name
	}

	insertSet := strings.Join(values, `, `)
	err := c.db.QueryRow(
		fmt.Sprintf(`INSERT INTO "categories" SET %s RETURNING "id", "name", "tags"`, insertSet),
		args...,
	).Scan(&category.Id, &category.Name, &category.Tags)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func (c DBCategories) Update(id int, categoryInput input.CategoryInput) (*Category, error) {
	category := Category{Id: id}
	values := make([]string, 0)
	args := make([]interface{}, 0)
	numberOfArguments := 0

	if categoryInput.Tags != nil {
		category.Tags = CategoryTags{}
		err := category.Tags.Scan(*categoryInput.Tags)
		if err != nil {
			return nil, err
		}
		numberOfArguments++
		values = append(values, fmt.Sprintf(`"tags"=$%d`, numberOfArguments))
		args = append(args, category.Tags)
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

	result, err := c.db.Exec(
		fmt.Sprintf(`UPDATE "categories" SET %s WHERE "id"=$%d RETURNING "id"`, updateSet, numberOfArguments),
		args...,
	)
	if result == nil {
		return nil, err
	}
	return &category, err
}

func (c DBCategories) Delete(id int) (bool, error) {
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
	category := Category{Id: categoryId}
	row := c.db.QueryRow(
		`SELECT "name", "tags" FROM "categories" WHERE "id"=$1`,
		categoryId,
	).Scan(
		&category.Name, &category.Tags,
	)

	if row == nil {
		return nil, errors.New(`category not found`)
	}
	if len(row.Error()) != 0 {
		return nil, errors.New(row.Error())
	}
	return &category, nil
}

func (categoryTags CategoryTags) Value() (driver.Value, error) {
	jsonCategoryTags, err := json.Marshal(categoryTags)
	return jsonCategoryTags, err
}

func (categoryTags *CategoryTags) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("failed to initialize variable for source")
	}
	err := json.Unmarshal(source, &categoryTags)
	if err != nil {
		return err
	}
	return nil
}
