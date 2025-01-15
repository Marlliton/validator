package rule

import (
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/Marlliton/validator/fail"
)

const (
	ErrMustBeAnString     = "the field '%s' must be an string, but received '%v'"
	ErrValidEmail         = "the field '%s' is ivalid. '%v'"
	ErrInvalidPhoneNumber = "the field '%s' is not a valid phone number %v"
	ErrNonEmptyString     = "the field '%s' must be a non-empty string representing a valid URL"
	ErrInvalidUrl         = "the field '%s' is not a valid URL: %s"
)

func String() Rule {
	return func(key string, value interface{}) *fail.Error {
		if _, ok := value.(string); !ok {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrMustBeAnString, key, reflect.TypeOf(value)),
			}
		}
		return nil
	}
}

func ValidEmail() Rule {
	return func(key string, value interface{}) *fail.Error {
		_, err := mail.ParseAddress(value.(string))
		if err != nil {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrValidEmail, key, err.Error()),
			}
		}
		return nil
	}
}

func ValidPhoneNumber() Rule {
	return func(key string, value interface{}) *fail.Error {
		e164Regex := `^\+[1-9]\d{3,14}$`
		re := regexp.MustCompile(e164Regex)
		phoneNumber := strings.ReplaceAll(value.(string), " ", "")

		isValid := re.Find([]byte(phoneNumber)) != nil
		if !isValid {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrInvalidPhoneNumber, key, value),
			}
		}

		return nil
	}
}

func ValidURL() Rule {
	return func(key string, value interface{}) *fail.Error {
		urlStr, ok := value.(string)
		if !ok || urlStr == "" {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrNonEmptyString, key),
			}
		}

		parsedURL, err := url.ParseRequestURI(urlStr)
		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			return &fail.Error{
				Field:   key,
				Message: fmt.Sprintf(ErrInvalidUrl, key, urlStr),
			}
		}

		return nil
	}
}
