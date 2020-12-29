package go3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo3(t *testing.T) {
	s := "abcabcbb"
	assert.Equal(t, 3, lengthOfLongestSubstring(s))
	s = "bbbbb"
	assert.Equal(t, 1, lengthOfLongestSubstring(s))
	s = "pwwkew"
	assert.Equal(t, 3, lengthOfLongestSubstring(s))
	s = ""
	assert.Equal(t, 0, lengthOfLongestSubstring(s))
}
