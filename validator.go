package validator

import (
	"reflect"

	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type validator struct {
	fieldRules map[string]rule.Rules
}

func New() *validator {
	return &validator{fieldRules: make(map[string]rule.Rules)}
}
func (v *validator) FieldRules() map[string]rule.Rules {
	return v.fieldRules
}

func (v *validator) Add(field string, rules rule.Rules) {
	if v.fieldRules == nil {
		v.fieldRules = make(map[string]rule.Rules)
	}

	v.fieldRules[field] = append(v.fieldRules[field], rules...)
}

func (v *validator) Validate(data interface{}) []*validator_error.ValidatorError {
	var errors []*validator_error.ValidatorError

	val := reflect.ValueOf(data)
	typeOfData := reflect.TypeOf(data)
	if val.Kind() != reflect.Struct {
		panic("Validate: o dado fornecido não é uma struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := typeOfData.Field(i)
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
