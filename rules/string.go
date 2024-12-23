package rules

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

func ValidEmail() Rule {
	return func(key string, value interface{}) error {
		_, err := mail.ParseAddress(value.(string))
		if err != nil {
			return fmt.Errorf("the field '%s' is ivalid. %v", key, err.Error())
		}
		return nil
	}
}

func ValidPhoneNumber() Rule {
	return func(key string, value interface{}) error {
		e164Regex := `^\+[1-9]\d{3,14}$`
		re := regexp.MustCompile(e164Regex)
		phoneNumber := strings.ReplaceAll(value.(string), " ", "")

		isValid := re.Find([]byte(phoneNumber)) != nil
		if !isValid {
			return fmt.Errorf("the field '%s' is not a valid phone number %v", key, value)
		}

		return nil
	}
}
