package input

import "errors"

type CategoryInput struct {
	Name *string   `json:"name"`
	Tags *[]string `json:"tags"`
}

func (ci *CategoryInput) Validate() error {
	if ci.Name == nil {
		return errors.New("category name can't be empty")
	}
	if ci.Tags == nil || len(*ci.Tags) == 0 {
		return errors.New("category tags can't be empty")
	}
	return nil
}
