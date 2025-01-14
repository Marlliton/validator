package rule

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/validator_error"
)

const (
	ErrMinLength                = "the field '%s' must be greater than %d"
	ErrMaxLength                = "the field '%s' must be less than %d"
	ErrCannotBeValidatedMin     = "the field '%s' cannot be validated with MinLength"
	ErrCannotBeValidatedMax     = "the field '%s' cannot be validated with MaxLength"
	ErrCannotBeValidatedExactly = "the field '%s' cannot be validated with ExactLength"
	ErrExactLength              = "the field '%s' must have exactly %d"
)

func MinLength(min int) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() < min {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf(ErrMinLength, key, min)}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf(ErrCannotBeValidatedMin, key),
			}
		}
		return nil
	}
}

func MaxLength(max int) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() > max {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf(ErrMaxLength, key, max),
				}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf(ErrCannotBeValidatedMax, key),
			}
		}
		return nil
	}
}

func ExactLength(length int) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() != length {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf(ErrExactLength, key, length),
				}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf(ErrCannotBeValidatedExactly, key),
			}
		}
		return nil
	}
}
