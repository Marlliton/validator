package rules

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/validator_error"
)

func MinLength(min int) Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() < min {
				return &validator_error.ValidatorError{
					Field:   key,
					Message: fmt.Sprintf("the field '%s' must be greater than %d", key, min)}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' cannot be validated with MinLength", key),
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
					Message: fmt.Sprintf("the field '%s' must be less than %d", key, max),
				}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' cannot be validated with MaxLength", key),
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
					Message: fmt.Sprintf("the field '%s' must have exactly %d", key, length),
				}
			}
		default:
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' cannot be validated with ExactLength", key),
			}
		}
		return nil
	}
}
