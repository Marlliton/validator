package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bool(t *testing.T) {
	rule := Bool()
	tests := []struct {
		key      string
		value    interface{}
		expected string
	}{
		{"boolean", true, ""},
		{"boolean", false, ""},

		{"boolean", 1234, "the field boolean is not a boolean value"},
		{"boolean", "1234", "the field boolean is not a boolean value"},
		{"boolean", nil, "the field boolean is not a boolean value"},
	}

	for _, tc := range tests {
		err := rule(tc.key, tc.value)
		if tc.expected == "" {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "the field boolean is not a boolean value")
		}
	}
}
