package rules

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/validator_error"
)

func Required() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		errMsg := &validator_error.ValidatorError{
			Field:   key,
			Message: fmt.Sprintf("the field '%s' is required", key),
		}

		if value == nil {
			return errMsg
		}

		val := reflect.ValueOf(value)

		switch val.Kind() {
		case reflect.String:
			if val.String() == "" {
				return errMsg
			}
		case reflect.Slice:
			if val.Len() == 0 {
				return errMsg
			}
		case reflect.Map:
			if len(val.MapKeys()) == 0 {
				return errMsg
			}
		default:
			if val.IsZero() {
				return errMsg
			}
		}

		return nil
	}
}
