package go4

import (
	"fmt"
	"strconv"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1, length2 := len(nums1), len(nums2)
	length := length1 + length2

	if length1 == 0 && length2 == 1 {
		return toFloat(float64(nums2[0]))
	}

	if length2 == 0 && length1 == 1 {
		return toFloat(float64(nums1[0]))
	}

	k, j := 0, 0
	for i := 0; i < length/2-1; i++ {
		if (j >= length2) || (k < length1 && nums1[k] < nums2[j]) {
			k++
		} else {
			j++
		}
	}

	result := make([]int, 0)

	for i := length/2 - 1; i < length/2+1; i++ {
		if (j >= length2) || (k < length1 && nums1[k] < nums2[j]) {
			result = append(result, nums1[k])
			k++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	if length%2 == 0 {
		return toFloat(float64(result[0]+result[1]) / 2)
	}
	return toFloat(float64(result[1]))
}

func toFloat(f float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", f), 64)
	return value
}
