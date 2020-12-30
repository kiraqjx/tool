package go8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo8(t *testing.T) {
	s := "42"
	assert.Equal(t, 42, myAtoi(s))
	s = "   -42"
	assert.Equal(t, -42, myAtoi(s))
	s = "4193 with words"
	assert.Equal(t, 4193, myAtoi(s))
	s = "words and 987"
	assert.Equal(t, 0, myAtoi(s))
	s = "-91283472332"
	assert.Equal(t, -2147483648, myAtoi(s))
	s = "+1"
	assert.Equal(t, 1, myAtoi(s))
	s = "9223372036854775808"
	assert.Equal(t, 2147483647, myAtoi(s))
}
