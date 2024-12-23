package rules

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/Marlliton/validator/validator_error"
)

func ValidEmail() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		_, err := mail.ParseAddress(value.(string))
		if err != nil {
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' is ivalid. %v", key, err.Error()),
			}
		}
		return nil
	}
}

func ValidPhoneNumber() Rule {
	return func(key string, value interface{}) *validator_error.ValidatorError {
		e164Regex := `^\+[1-9]\d{3,14}$`
		re := regexp.MustCompile(e164Regex)
		phoneNumber := strings.ReplaceAll(value.(string), " ", "")

		isValid := re.Find([]byte(phoneNumber)) != nil
		if !isValid {
			return &validator_error.ValidatorError{
				Field:   key,
				Message: fmt.Sprintf("the field '%s' is not a valid phone number %v", key, value),
			}
		}

		return nil
	}
}
