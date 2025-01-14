package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	t.Run("Verify if value is an string", func(t *testing.T) {
		rule := String()
		err := rule("test", "string")
		assert.Nil(t, err)
	})

	t.Run("return error if value is not string", func(t *testing.T) {
		rule := String()
		err := rule("test", 15.5)
		assert.NotNil(t, err)

		err = rule("test", "15.5")
		assert.Nil(t, err)
	})
}

func Test_ValidURL(t *testing.T) {
	rule := ValidURL()

	tests := []struct {
		key      string
		value    interface{}
		expected string
	}{
		// URLs válidas
		{"url", "https://valid-url.com", ""},
		{"url", "http://example.org", ""},
		{"url", "https://sub.domain.com", ""},
		{"url", "https://www.google.com/search?q=golang", ""},
		{"url", "ftp://ftp.example.com/file.txt", ""},
		{"url", "http://localhost:8080", ""},

		// URLs inválidas
		{"url", "", "the field url is not a valid URL"},
		{"url", "invalid-url", "the field url is not a valid URL"},
		{"url", "https://", "the field url is not a valid URL"},
		{"url", "ftp://", "the field url is not a valid URL"},

		// Casos com outros tipos
		{"url", 12345, "the field url is not a valid URL"},
		{"url", true, "the field url is not a valid URL"},
		{"url", nil, "the field url is not a valid URL"},
	}

	for _, tc := range tests {
		err := rule(tc.key, tc.value)
		if tc.expected == "" {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err, "expected an error for value: %v", tc.value)
		}
	}
}

func Test_ValidEmail(t *testing.T) {
	rule := ValidEmail()

	tests := []struct {
		key      string
		value    interface{}
		expected string
	}{
		{"email", "valid@example.com", ""},
		{"email", "invalid-email", "the field 'email' is ivalid. 'mail: missing '@' or angle-addr'"},
		{"email", "", "the field 'email' is ivalid. 'mail: no address'"},
		{"email", "test@.com", "the field 'email' is ivalid. 'mail: missing '@' or angle-addr'"},
		{"email", "test@domain..com", "the field 'email' is ivalid. 'mail: missing '@' or angle-addr'"},
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
		{"phone", "+55 79 9 9912 9999", ""},
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
