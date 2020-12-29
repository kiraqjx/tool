package go4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo4(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	assert.Equal(t, 2.00000, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	assert.Equal(t, 2.50000, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	assert.Equal(t, 0.00000, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{}
	nums2 = []int{1}
	assert.Equal(t, 1.00000, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{2}
	nums2 = []int{}
	assert.Equal(t, 2.00000, findMedianSortedArrays(nums1, nums2))
}
