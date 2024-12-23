package validator_error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidatorError_Error(t *testing.T) {
	t.Run("Test ValidatorError creation", func(t *testing.T) {
		ve := &ValidatorError{
			Field:   "username",
			Message: "must be at least 3 characters long",
		}

		assert.Equal(t, "username", ve.Field)
		assert.Equal(t, "must be at least 3 characters long", ve.Message)
	})

	t.Run("Test Error method", func(t *testing.T) {
		ve := &ValidatorError{
			Field:   "email",
			Message: "must be a valid email address",
		}

		expectedErrorMessage := "must be a valid email address"
		assert.Equal(t, expectedErrorMessage, ve.Error())
	})

	t.Run("Test Error method with different values", func(t *testing.T) {
		ve := &ValidatorError{
			Field:   "password",
			Message: "must contain at least one uppercase letter",
		}

		expectedErrorMessage := "must contain at least one uppercase letter"
		assert.Equal(t, expectedErrorMessage, ve.Error())
	})
}
