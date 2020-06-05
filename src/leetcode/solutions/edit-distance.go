package solutions

import "leetcode/lib"

func MinDistance(word1 string, word2 string) int {
	distanceBetweenPrefixes := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		distanceBetweenPrefixes[i] = make([]int, len(word2))
		for j, _ := range distanceBetweenPrefixes[i] {
			distanceBetweenPrefixes[i][j] = -1
		}
	}

	return computeDistanceBetweenPrefixes(word1, len(word1)-1, word2, len(word2)-1, distanceBetweenPrefixes)
}

func computeDistanceBetweenPrefixes(word1 string, i int, word2 string, j int, prefixes [][]int) int {
	if i < 0 {
		return j + 1
	} else if j < 0 {
		return i + 1
	}

	if prefixes[i][j] == -1 {
		if word1[i] == word2[j] {
			prefixes[i][j] = computeDistanceBetweenPrefixes(word1, i-1, word2, j-1, prefixes)
		} else {
			substituteLast := computeDistanceBetweenPrefixes(word1, i-1, word2, j-1, prefixes)
			addLast := computeDistanceBetweenPrefixes(word1, i, word2, j-1, prefixes)
			deleteLast := computeDistanceBetweenPrefixes(word1, i-1, word2, j, prefixes)
			prefixes[i][j] = 1 + lib.MinInt(substituteLast, lib.MinInt(addLast, deleteLast))
		}
	}
	return prefixes[i][j]
}
