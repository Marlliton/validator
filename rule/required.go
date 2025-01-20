package rule

import (
	"fmt"
	"reflect"

	"github.com/Marlliton/validator/fail"
)

const ErrRequired = "the field '%s' is required"

func Required() Rule {
	return func(key string, value interface{}) *fail.Error {
		errMsg := &fail.Error{
			Field:   key,
			Message: fmt.Sprintf(ErrRequired, key),
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
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return nil

		case reflect.Bool:
			return nil
		default:
			if val.IsZero() {
				return errMsg
			}
		}

		return nil
	}
}
