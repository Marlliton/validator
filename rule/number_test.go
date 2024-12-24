package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinValue(t *testing.T) {
	t.Run("Value equal to minimum", func(t *testing.T) {
		rule := MinValue(10)
		err := rule("test", 10)
		assert.Nil(t, err)
	})

	t.Run("Value greater than minimum", func(t *testing.T) {
		rule := MinValue(10)
		err := rule("test", 15)
		assert.Nil(t, err)
	})

	t.Run("Value less than minimum", func(t *testing.T) {
		rule := MinValue(10)
		err := rule("test", 5)
		assert.NotNil(t, err)
	})
}

func Test_MaxValue(t *testing.T) {
	t.Run("Value equal to maximum", func(t *testing.T) {
		rule := MaxValue(10)
		err := rule("test", 10)
		assert.Nil(t, err)
	})

	t.Run("Value greater than maximum", func(t *testing.T) {
		rule := MaxValue(10)
		err := rule("test", 15)
		assert.NotNil(t, err)
	})

	t.Run("Value less than maximum", func(t *testing.T) {
		rule := MaxValue(10)
		err := rule("test", 5)
		assert.Nil(t, err)
	})
}
