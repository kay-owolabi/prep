package solutions

import "sort"

type cost [][]int

func (c cost) Len() int { return len(c) }

func (c cost) Less(i, j int) bool {
	return (c[i][0] - c[i][1]) < (c[j][0] - c[j][1])
}
func (c cost) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func TwoCitySchedCost(costs [][]int) int {
	var allCosts cost = costs
	sort.Sort(allCosts)
	n := len(allCosts) / 2
	var res int
	for i := 0; i < n; i++ {
		res += allCosts[i][0] + allCosts[i+n][1]
	}
	return res
}
