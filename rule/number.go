package rule

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/fail"
)

const (
	ErrMustBeAnInteger = "the field '%s' must be an integer, but received %v"
	ErrMinValue        = "the field '%s' must be grater than or equal to %v"
	ErrMaxValue        = "the field '%s' must be less than or equal to %v"
)

func Int() Rule {
	return func(key string, value interface{}) *fail.Error {
		if _, ok := value.(int); !ok {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeAnInteger, key, reflect.TypeOf(value)),
			}
		}
		return nil
	}
}

func MinValue[T int | float64](min T) Rule {
	return func(key string, value interface{}) *fail.Error {
		switch v := value.(type) {
		case T:
			if v < min {
				return &fail.Error{
					Field:   key,
					Message: fmt.Sprintf(ErrMinValue, key, min),
				}
			}

		default:
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeAnInteger, key, reflect.TypeOf(value)),
			}
		}

		return nil
	}

}

func MaxValue[T int | float64](max T) Rule {
	return func(key string, value interface{}) *fail.Error {
		switch v := value.(type) {
		case T:
			if v > max {
				return &fail.Error{
					Field:   key,
					Message: fmt.Sprintf(ErrMaxValue, key, max),
				}
			}
		default:
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeAnInteger, key, reflect.TypeOf(value)),
			}
		}

		return nil
	}
}
