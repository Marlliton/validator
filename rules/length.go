package rules

import (
	"fmt"
	"reflect"
)

func MinLength(min int) Rule {
	return func(key string, value interface{}) error {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() < min {
				return fmt.Errorf("o campo %s deve ter no mínimo %d", key, min)
			}
		default:
			return fmt.Errorf("o campo '%s' não pode ser validado com MinLength", key)
		}
		return nil
	}
}

func MaxLength(max int) Rule {
	return func(key string, value interface{}) error {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() > max {
				return fmt.Errorf("o campo %s deve ter no máximo %d", key, max)
			}
		default:
			return fmt.Errorf("o campo '%s' não pode ser validado com MaxLength", key)
		}
		return nil
	}
}

func ExactLength(length int) Rule {
	return func(key string, value interface{}) error {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			if v.Len() != length {
				return fmt.Errorf("o campo %s deve ter exatamente %d", key, length)
			}
		default:
			return fmt.Errorf("o campo '%s' não pode ser validado com MaxLength", key)
		}
		return nil
	}
}
