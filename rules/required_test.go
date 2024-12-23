package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	required := Required()

	t.Run("should return error for nil value", func(t *testing.T) {
		err := required("field", nil)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})

	t.Run("should return error for empty string", func(t *testing.T) {
		err := required("field", "")
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})

	t.Run("should return error for empty slice", func(t *testing.T) {
		err := required("field", []int{})
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})

	t.Run("should return error for empty map", func(t *testing.T) {
		err := required("field", map[string]string{})
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})

	t.Run("should not return error for non-empty string", func(t *testing.T) {
		err := required("field", "value")
		assert.Nil(t, err)
	})

	t.Run("should not return error for non-empty slice", func(t *testing.T) {
		err := required("field", []int{1, 2, 3})
		assert.Nil(t, err)
	})

	t.Run("should not return error for non-empty map", func(t *testing.T) {
		err := required("field", map[string]string{"key": "value"})
		assert.Nil(t, err)
	})

	t.Run("should not return error for non-zero integer", func(t *testing.T) {
		err := required("field", 1)
		assert.Nil(t, err)
	})

	t.Run("should return error for zero integer", func(t *testing.T) {
		err := required("field", 0)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})
}
