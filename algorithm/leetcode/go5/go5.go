package go5

//bab
func longestPalindrome(s string) string {
	all := make(map[int]string)
	strLen := len(s)
	max := 0
	for index := range s {
		i := 1
		needBreak1 := true
		needBreak2 := true
		for {
			if needBreak1 {
				if index-i < 0 || index+i >= strLen || s[index-i] != s[index+i] {
					if 2*i-1 > max {
						max = 2*i - 1
					}
					all[2*i-1] = s[index-i+1 : index+i]
					needBreak1 = false
				}
			}
			if needBreak2 {
				if index-i < 0 || index+i > strLen || s[index-i] != s[index+i-1] {
					if 2*i-2 > max {
						max = 2*i - 2
					}
					if index-i+1 != index+i-1 {
						all[2*i-2] = s[index-i+1 : index+i-1]
					}
					needBreak2 = false
				}
			}
			if !needBreak1 && !needBreak2 {
				break
			}
			i++
		}
	}
	return all[max]
}
