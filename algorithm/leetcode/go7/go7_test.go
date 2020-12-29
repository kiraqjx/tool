package go7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo7(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 21, reverse(120))
}
