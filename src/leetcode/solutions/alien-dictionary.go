package solutions

import "leetcode/lib"

func isAlienSorted(words []string, order string) bool {
	dictionary := map[byte]int{}

	for i, letter := range order {
		dictionary[byte(letter)] = i
	}

	for i := 0; i < (len(words) - 1); i++ {

		maxLen := lib.MinInt(len(words[i]), len(words[i+1])) - 1

		for j := 0; j <= maxLen; j++ {
			weight0 := dictionary[words[i][j]]
			weight1 := dictionary[words[i+1][j]]
			if weight0 < weight1 {
				break
			}

			if weight0 > weight1 {
				return false
			}
		}

		if words[i][maxLen] == words[i+1][maxLen] && len(words[i]) > len(words[i+1]) {
			return false
		}
	}
	return true
}
