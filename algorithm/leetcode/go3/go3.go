package go3

func lengthOfLongestSubstring(s string) int {
	group := make(map[rune]int)
	max := 0
	for index, char := range s {
		if index2, ok := group[char]; ok {
			if len(group) > max {
				max = len(group)
			}
			for key, value := range group {
				if value <= index2 {
					delete(group, key)
				}
			}
		}
		group[char] = index
	}
	if len(group) > max {
		max = len(group)
	}
	return max
}
