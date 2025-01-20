package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Required(t *testing.T) {
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

	t.Run("should not return error for zero integer", func(t *testing.T) {
		err := required("field", 0)
		assert.Nil(t, err)
	})

	t.Run("should not return error for zero float", func(t *testing.T) {
		err := required("field", 0.0)
		assert.Nil(t, err)
	})

	t.Run("should not return error for non-zero float", func(t *testing.T) {
		err := required("field", 1.23)
		assert.Nil(t, err)
	})

	t.Run("should not return error for true boolean", func(t *testing.T) {
		err := required("field", true)
		assert.Nil(t, err)
	})

	t.Run("should not return error for false boolean", func(t *testing.T) {
		err := required("field", false)
		assert.Nil(t, err)
	})

	t.Run("should return error for empty struct", func(t *testing.T) {
		err := required("field", struct{}{})
		assert.NotNil(t, err)
	})

	t.Run("should not return error for non-empty struct", func(t *testing.T) {
		err := required("field", struct{ Name string }{Name: "Test"})
		assert.Nil(t, err)
	})

	t.Run("should return error for nil pointer", func(t *testing.T) {
		var ptr *int
		err := required("field", ptr)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "the field 'field' is required")
	})

	t.Run("should not return error for non-nil pointer", func(t *testing.T) {
		ptr := 42
		err := required("field", &ptr)
		assert.Nil(t, err)
	})
}
