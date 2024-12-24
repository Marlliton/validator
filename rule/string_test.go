package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidEmail(t *testing.T) {
	rule := ValidEmail()

	tests := []struct {
		key      string
		value    interface{}
		expected string
	}{
		{"email", "valid@example.com", ""},
		{"email", "invalid-email", "the field 'email' is ivalid. mail: missing '@' or angle-addr"},
		{"email", "", "the field 'email' is ivalid. mail: no address"},
		{"email", "test@.com", "the field 'email' is ivalid. mail: missing '@' or angle-addr"},
		{"email", "test@domain..com", "the field 'email' is ivalid. mail: missing '@' or angle-addr"},
	}

	for _, test := range tests {
		err := rule(test.key, test.value)
		if test.expected == "" {
			assert.Nil(t, err, "expected no error for value: %v", test.value)
		} else {
			assert.NotNil(t, err, "expected an error for value: %v", test.value)
			assert.Contains(t, err.Error(), test.expected)
		}
	}
}

func Test_ValidPhoneNumber(t *testing.T) {
	rule := ValidPhoneNumber()

	tests := []struct {
		key      string
		value    interface{}
		expected string
	}{
		{"phone", "+1234567890", ""},
		{"phone", "+1 234 567 890", ""},
		{"phone", "+123", "the field 'phone' is not a valid phone number"},
		{"phone", "1234567890", "the field 'phone' is not a valid phone number"},
		{"phone", "+12 345678901234567890", "the field 'phone' is not a valid phone number"},
		{"phone", "", "the field 'phone' is not a valid phone number"},
	}

	for _, test := range tests {
		err := rule(test.key, test.value)
		if test.expected == "" {
			assert.Nil(t, err, "expected no error for value: %v", test.value)
		} else {
			assert.NotNil(t, err, "expected an error for value: %v", test.value)
			assert.Contains(t, err.Error(), test.expected)
		}
	}
}
