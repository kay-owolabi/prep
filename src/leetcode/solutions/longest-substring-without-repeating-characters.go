package solutions

func LengthOfLongestSubstring(s string) int {
	var start, maxSize int
	charMap := make(map[byte]int)

	checkMax := func(maxSize *int, charMap map[byte]int) {
		if *maxSize < len(charMap) {
			*maxSize = len(charMap)
		}
	}

	for i, char := range s {
		char := byte(char)
		index, ok := charMap[char]
		if ok {
			checkMax(&maxSize, charMap)
			for ; index != start; start++ {
				delete(charMap, s[start])
			}
			start++
		}
		charMap[char] = i
	}
	checkMax(&maxSize, charMap)
	return maxSize
}
