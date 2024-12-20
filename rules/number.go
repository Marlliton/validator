package rules

import (
	"fmt"
	"reflect"
)

func MinValue[T int | float64](min T) Rule {
	return func(key string, value interface{}) error {
		switch v := value.(type) {
		case T:
			if v < min {
				return fmt.Errorf("the field '%s' must be grater than or equal to %v", key, min)
			}

		default:
			return fmt.Errorf("the field '%s' must be a number of type %v", key, reflect.TypeOf(min))
		}

		return nil
	}

}

func MaxValue[T int | float64](max T) Rule {
	return func(key string, value interface{}) error {
		switch v := value.(type) {
		case T:
			if v > max {
				return fmt.Errorf("the field '%s' must be less than or equal to %v", key, max)
			}
		default:
			return fmt.Errorf("the field '%s' must be a number of type %v", key, reflect.TypeOf(max))
		}

		return nil
	}
}
