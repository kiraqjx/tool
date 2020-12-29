package go6

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	strLen := len(s)
	groupNum := 2*numRows - 2
	group := strLen / groupNum
	convertS := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= group; j++ {
			if j*groupNum+i < strLen {
				convertS = append(convertS, s[j*groupNum+i])
			}
			if i != 0 && i != numRows-1 {
				if (j+1)*groupNum-i < strLen {
					convertS = append(convertS, s[(j+1)*groupNum-i])
				}
			}
		}
	}
	return string(convertS)
}
