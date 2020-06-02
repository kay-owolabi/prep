package solutions

import "leetcode/lib"

func NumEquivDominoPairs(dominoes [][]int) int {
	hashMap := map[int]int{}
	result := 0
	for i := 0; i < len(dominoes); i = i + 1 {
		x := lib.MinInt(dominoes[i][0], dominoes[i][1])
		y := lib.MaxInt(dominoes[i][0], dominoes[i][1])
		key := x*10 + y
		value, ok := hashMap[key]
		if !ok {
			value = 0
		}
		result += value
		hashMap[key] = value + 1
	}
	return result
}
