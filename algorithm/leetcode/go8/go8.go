package go8

func myAtoi(s string) int {
	result := 0
	isMinus := false
	i := 0
	strLen := len(s)
	for {
		if i >= strLen {
			return 0
		} else if s[i] == 32 {
			i++
			continue
		} else if string(s[i]) == "-" {
			isMinus = true
			i++
			break
		} else if string(s[i]) == "+" {
			i++
			break
		} else if s[i] < 48 || s[i] > 57 {
			return 0
		}
		result = result*10 + int(s[i]-48)
		i++
		break
	}
	for {
		if i >= strLen {
			break
		} else if s[i] == 32 {
			break
		} else if s[i] < 48 || s[i] > 57 {
			break
		}
		result = result*10 + int(s[i]-48)
		if !isMinus && result > 2147483647 {
			result = 2147483647
			break
		} else if isMinus && result > 2147483648 {
			result = 2147483648
			break
		}
		i++
	}

	if isMinus && result != 0 {
		return -result
	}
	return result
}
