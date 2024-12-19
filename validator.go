package validator

import (
	"fmt"
	"reflect"
)

type Rule func(key string, value interface{}) error

type Rules []Rule

type Validator struct {
	fieldRules map[string]Rules
}

func (v *Validator) Add(field string, rules Rules) {
	if v.fieldRules == nil {
		v.fieldRules = make(map[string]Rules)
	}

	v.fieldRules[field] = append(v.fieldRules[field], rules...)
}

func (v *Validator) Validate(data map[string]interface{}) []error {
	var errors []error
	for field, rules := range v.fieldRules {
		value, exist := data[field]
		if !exist {
			return nil
		}

		for _, rule := range rules {
			if err := rule(field, value); err != nil {
				errors = append(errors, err)
			}
		}
	}

	return errors
}

func MinLength(min int) Rule {
	return func(key string, value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("o campo '%s' deve ser uma string", key)
		}
		if len(str) < min {
			return fmt.Errorf("o campo '%s' deve ter no mínimo %d caracteres", key, min)
		}
		return nil
	}
}

func LengthBetween(start, end int, key string, value interface{}) Rule {
	return func(key string, value interface{}) error {
		str, ok := value.(string)
		if ok {
			if len(str) < start || len(str) > end {
				return fmt.Errorf("%s must be between %d and %d", key, start, end)
			}
			return nil
		}

		if isArrayOrSlice(value) {
			length := reflect.ValueOf(value).Len()
			if length < start || length > end {
				return fmt.Errorf("%s must be between %d and %d items", key, start, end)
			}
			return nil
		}
		return fmt.Errorf("%s must be a string, array, or slice", key)
	}

}

func isArrayOrSlice(input interface{}) bool {
	valueType := reflect.TypeOf(input)
	if valueType == nil {
		return false
	}

	switch valueType.Kind() {
	case reflect.Array:
		return true
	case reflect.Slice:
		return true
	default:
		return false
	}
}
