package go9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo9(t *testing.T) {
	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, false, isPalindrome(-121))
	assert.Equal(t, false, isPalindrome(10))
	assert.Equal(t, true, isPalindrome(0))
}
