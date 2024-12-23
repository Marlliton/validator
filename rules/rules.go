package rules

import "github.com/Marlliton/validator/validator_error"

type Rule func(key string, value interface{}) *validator_error.ValidatorError

type Rules []Rule
