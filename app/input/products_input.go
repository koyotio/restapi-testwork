package input

import "errors"

type ProductInput struct {
	CategoryId *int    `json:"category_id"`
	Articul    *string `json:"articul"`
	Name       *string `json:"name"`
}

func (pi *ProductInput) Validate() error {
	if pi.CategoryId == nil {
		return errors.New("product category_id can't be empty")
	}
	if pi.Articul == nil {
		return errors.New("product articul can't be empty")
	}
	if pi.Name == nil {
		return errors.New("product name can't be empty")
	}
	return nil
}
