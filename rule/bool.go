package rule

import (
	"fmt"

	"github.com/Marlliton/validator/validator_error"
)

const ErrMustBeABool = "the field '%s' must be a boolean"

func Bool() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		if _, ok := value.(bool); !ok {
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeABool, key),
			}
		}
		return nil
	}
}
