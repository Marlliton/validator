package rule

import "github.com/Marlliton/validator/fail"

type Rule func(key string, value interface{}) *fail.Error

type Rules []Rule
