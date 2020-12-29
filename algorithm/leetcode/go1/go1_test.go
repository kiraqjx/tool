package go1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	assert.Equal(t, []int{0, 1}, twoSum(nums, 9))
}
