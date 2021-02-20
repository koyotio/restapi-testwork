package input

import "errors"

type CategoryInput struct {
	Name *string `json:"name"`
	Tags *string `json:"tags"`
}

func (ci CategoryInput) Validate() error {
	if ci.Name == nil && (ci.Tags == nil || len(*ci.Tags) == 0) {
		return errors.New("input can't be empty")
	}
	return nil
}
