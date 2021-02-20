package input

import "errors"

type ProductInput struct {
	CategoryId *int    `json:"category_id"`
	Articul    *string `json:"articul"`
	Name       *string `json:"name"`
}

func (pi ProductInput) Validate() error {
	if pi.CategoryId == nil && pi.Articul == nil && pi.Name == nil {
		return errors.New("input can't be empty")
	}
	return nil
}
