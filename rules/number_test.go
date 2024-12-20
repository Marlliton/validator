package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinValue(t *testing.T) {
	rule := MinValue(10)
	err := rule("test", "d")
	assert.Nil(t, err)
}
