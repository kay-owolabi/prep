package solutions

import "leetcode/lib"

func FindMaxLength(nums []int) int {
	countMap := map[int]int{
		0: -1,
	}
	maxLen, count := 0, 0

	for i, val := range nums {
		if val == 0 {
			count -= 1
		} else {
			count += 1
		}

		index, ok := countMap[count]

		if ok {
			maxLen = lib.MaxInt(maxLen, i-index)
		} else {
			countMap[count] = i
		}
	}
	return maxLen
}
