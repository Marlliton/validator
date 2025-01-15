package validator

import (
	"reflect"

	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
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

func (v *validator) Validate(data interface{}) []*fail.Error {
	var errors []*fail.Error

	val := reflect.ValueOf(data)

	v.validateStruct(val, "", &errors)

	return errors
}

func (v *validator) validateStruct(val reflect.Value, prefix string, errors *[]*fail.Error) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		panic("Validate: o dado fornecido não é uma struct")
	}

	typeOfData := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typeOfData.Field(i)
		fieldValue := val.Field(i)

		fieldName := field.Name
		fullFieldName := fieldName

		if prefix != "" {
			fullFieldName = prefix + "." + fieldName
		}

		if fieldValue.Kind() == reflect.Struct {
			v.validateStruct(fieldValue, fullFieldName, errors)
			continue
		}

		if rules, exists := v.fieldRules[fullFieldName]; exists {
			for _, rule := range rules {
				if err := rule(fullFieldName, fieldValue.Interface()); err != nil {
					*errors = append(*errors, err)
				}
			}
		}
	}
}
