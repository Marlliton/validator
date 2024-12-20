package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinLength(t *testing.T) {
	rule := MinLength(10)
	errStr := rule("valid", "This is a valid text")
	assert.Nil(t, errStr)

	errStr = rule("invalid", "short")
	assert.NotNil(t, errStr)

	errArray := rule("array_valid", [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11})
	assert.Nil(t, errArray)

	errArray = rule("array_invalid", [1]int{1})
	assert.NotNil(t, errArray)

	errSlice := rule("slice_valid", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11})
	assert.Nil(t, errSlice)

	errSlice = rule("slice_invalid", []int{1, 2, 3})
	assert.NotNil(t, errSlice)
}

func Test_MaxLength(t *testing.T) {
	rule := MaxLength(10)
	err := rule("invalid", "This is a very long valid text")
	assert.NotNil(t, err)

	err = rule("valid", "short")
	assert.Nil(t, err)

	err = rule("exact", "1234567890")
	assert.Nil(t, err)

	err = rule("empty", "")
	assert.Nil(t, err)

	err = rule("valid_array", [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	assert.Nil(t, err)

	err = rule("invalid_array", [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11})
	assert.NotNil(t, err)

	errSlice := rule("slice_valid", []int{1, 2, 3})
	assert.Nil(t, errSlice)

	errSlice = rule("slice_invalid_long", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	assert.NotNil(t, errSlice)
}

func Test_ExactLength(t *testing.T) {
	rule := ExactLength(4)
	err := rule("invalid1", "ana")
	assert.NotNil(t, err)
	err = rule("invalid2", "John Doe")
	assert.NotNil(t, err)
	err = rule("invalid3", [1]int{1})
	assert.NotNil(t, err)
	err = rule("invalid4", map[string]int{"1": 1})
	assert.NotNil(t, err)

	err = rule("valid1", "1234")
	assert.Nil(t, err)
	err = rule("valid2", [4]int{1, 2, 3, 4})
	assert.Nil(t, err)
	err = rule("valid3", map[string]int{"1": 1, "2": 2, "3": 3, "4": 4})
	assert.Nil(t, err)
}
