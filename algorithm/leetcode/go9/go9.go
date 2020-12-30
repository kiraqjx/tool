package go9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	arr := make([]int, 0)
	for {
		arr = append(arr, x%10)
		x = x / 10
		if x == 0 {
			break
		}
	}

	lenArr := len(arr)
	for i := 0; i < lenArr/2; i++ {
		if arr[i] != arr[lenArr-1-i] {
			return false
		}
	}
	return true
}
