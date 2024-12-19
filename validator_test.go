package validator

import (
	"testing"
)

func TestValidatorAdd(t *testing.T) {
	val := Validator{}
	val.Add("Name", Rules{
		MinLength(10),
	})

	user := map[string]interface{}{
		"Name": "John aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}
	errors := val.Validate(user)
	t.Log(errors)
}
