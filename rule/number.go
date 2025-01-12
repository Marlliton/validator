package rule

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/validator_error"
)

func Int() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		if _, ok := value.(int); !ok {
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' must be an integer, but received %v", key, reflect.TypeOf(value)),
			}
		}
		return nil
	}
}

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
				Message: fmt.Sprintf("the field '%s' must be a number of type %v, but received %v", key, reflect.TypeOf(min), reflect.TypeOf(value)),
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
				Message: fmt.Sprintf("the field '%s' must be a number of type %v, but received %v", key, reflect.TypeOf(max), reflect.TypeOf(value)),
			}
		}

		return nil
	}
}
