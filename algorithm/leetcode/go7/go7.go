package go7

import "math"

func reverse(x int) int {
	isMinus := false
	if x < 0 {
		isMinus = true
	}
	arr := make([]int, 0)
	for {
		arr = append(arr, x%10)
		x = x / 10
		if x == 0 {
			break
		}
	}
	needX := 0
	arrLen := len(arr)
	for i := arrLen - 1; i >= 0; i-- {
		needX += arr[arrLen-1-i] * int(math.Pow(10, float64(i)))
	}
	if isMinus && needX > 0 {
		needX = -needX
	}
	if needX > int(math.Pow(2, 31)-1) || needX < int(-math.Pow(2, 31)) {
		return 0
	}
	return needX
}
