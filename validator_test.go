package validator

import (
	"testing"

	"github.com/Marlliton/validator/rules"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string
	Email string
	Age   int
}

func Test_New(t *testing.T) {
	t.Run("should create a new validator instance", func(t *testing.T) {
		v := New()
		assert.NotNil(t, v)
		assert.NotNil(t, v.FieldRules())
	})
}

func Test_Add(t *testing.T) {
	t.Run("should add rules for a field", func(t *testing.T) {
		v := New()
		v.Add("Name", rules.Rules{
			rules.Required(),
		})

		assert.NotNil(t, v.fieldRules)
		assert.Equal(t, 1, len(v.FieldRules()["Name"]))
	})
}

func Test_Validate(t *testing.T) {
	t.Run("should validate struct and return errors for invalid fields", func(t *testing.T) {
		v := New()
		v.Add("Name", rules.Rules{rules.Required()})
		v.Add("Email", rules.Rules{rules.Required()})

		data := TestStruct{
			Name:  "",
			Email: "",
		}

		errs := v.Validate(data)
		assert.Equal(t, 2, len(errs))
		assert.EqualError(t, errs[0], "the field 'Name' is required")
		assert.EqualError(t, errs[1], "the field 'Email' is required")
	})

	t.Run("should not return errors for a valid fields", func(t *testing.T) {
		v := New()
		v.Add("Name", rules.Rules{rules.Required()})
		v.Add("Email", rules.Rules{rules.Required()})

		data := TestStruct{
			Name:  "Valid Name",
			Email: "valid@zmail.com",
		}

		errs := v.Validate(data)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("should panic if is not a struct", func(t *testing.T) {
		v := New()
		assert.PanicsWithValue(t, "Validate: o dado fornecido não é uma struct", func() {
			v.Validate("not a struct")
		})
	})
}
