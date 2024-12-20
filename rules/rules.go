package rules

type Rule func(key string, value interface{}) error

type Rules []Rule
