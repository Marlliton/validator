package validator

import (
	"reflect"

	"github.com/Marlliton/validator/rules"
)

type validator struct {
	fieldRules map[string]rules.Rules
}

func New() *validator {
	return &validator{fieldRules: make(map[string]rules.Rules)}
}
func (v *validator) FieldRules() map[string]rules.Rules {
	return v.fieldRules
}

func (v *validator) Add(field string, r rules.Rules) {
	if v.fieldRules == nil {
		v.fieldRules = make(map[string]rules.Rules)
	}

	v.fieldRules[field] = append(v.fieldRules[field], r...)
}

func (v *validator) Validate(data interface{}) []error {
	var errors []error

	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	if val.Kind() != reflect.Struct {
		panic("Validate: o dado fornecido não é uma struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldValue := val.Field(i).Interface()

		if rules, exists := v.fieldRules[fieldName]; exists {
			for _, rule := range rules {
				if err := rule(fieldName, fieldValue); err != nil {
					errors = append(errors, err)
				}
			}
		}
	}

	return errors
}
