package go5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo5(t *testing.T) {
	s := "babad"
	assert.Equal(t, "aba", longestPalindrome(s))

	s = "cbbd"
	assert.Equal(t, "bb", longestPalindrome(s))

	s = "ac"
	assert.Equal(t, "c", longestPalindrome(s))

	s = "bb"
	assert.Equal(t, "bb", longestPalindrome(s))

	s = "ccc"
	assert.Equal(t, "ccc", longestPalindrome(s))
}
