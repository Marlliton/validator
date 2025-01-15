package rule

import (
	"fmt"

	"github.com/Marlliton/validator/fail"
)

const ErrMustBeABool = "the field '%s' must be a boolean"

func Bool() Rule {
	return func(key string, value interface{}) *fail.Error {
		if _, ok := value.(bool); !ok {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeABool, key),
			}
		}
		return nil
	}
}
