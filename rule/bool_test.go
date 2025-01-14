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

		{"boolean", 1234, "the field 'boolean' must be a boolean"},
		{"boolean", "1234", "the field 'boolean' must be a boolean"},
		{"boolean", nil, "the field 'boolean' must be a boolean"},
	}

	for _, tc := range tests {
		err := rule(tc.key, tc.value)
		if tc.expected == "" {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "the field 'boolean' must be a boolean")
		}
	}
}
