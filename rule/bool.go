package rule

import (
	"fmt"

	"github.com/Marlliton/validator/validator_error"
)

func Bool() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		if _, ok := value.(bool); !ok {
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field %s is not a boolean value", key),
			}
		}
		return nil
	}
}
