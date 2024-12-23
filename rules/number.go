package rules

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/validator_error"
)

func MinValue[T int | float64](min T) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		switch v := value.(type) {
		case T:
			if v < min {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf("the field '%s' must be grater than or equal to %v", key, min),
				}
			}

		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' must be a number of type %v", key, reflect.TypeOf(min)),
			}
		}

		return nil
	}

}

func MaxValue[T int | float64](max T) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		switch v := value.(type) {
		case T:
			if v > max {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf("the field '%s' must be less than or equal to %v", key, max),
				}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' must be a number of type %v", key, reflect.TypeOf(max)),
			}
		}

		return nil
	}
}
