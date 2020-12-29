package go6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo6(t *testing.T) {
	s := "LEETCODEISHIRING"
	numRows := 3
	assert.Equal(t, "LCIRETOESIIGEDHN", convert(s, numRows))

	s = "LEETCODEISHIRING"
	numRows = 4
	assert.Equal(t, "LDREOEIIECIHNTSG", convert(s, numRows))

	s = "ABC"
	numRows = 2
	assert.Equal(t, "ACB", convert(s, numRows))

}
