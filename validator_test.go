package validator

import (
	"testing"

	"github.com/Marlliton/validator/rules"
	"github.com/stretchr/testify/assert"
)

func Test_ValidatorAdd(t *testing.T) {
	val := validator{}
	val.Add("Name", rules.Rules{
		rules.MinLength(10),
		rules.MaxLength(20),
	})

	assert.Len(t, val.FieldRules()["Name"], 2)
}
