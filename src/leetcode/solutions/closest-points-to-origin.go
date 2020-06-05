package solutions

import (
	"container/heap"
)

type Coor [][]int

func (c Coor) Len() int {
	return len(c)
}

func (c Coor) Less(i, j int) bool {
	dist := func(point []int) int {
		return point[0]*point[0] + point[1]*point[1]
	}
	return dist(c[i]) < dist(c[j])
}

func (c Coor) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *Coor) Push(x interface{}) {
	*c = append(*c, x.([]int))
}

func (c *Coor) Pop() interface{} {
	var x []int
	x, *c = (*c)[len(*c)-1], (*c)[:len(*c)-1]
	return x
}

func KClosest(points [][]int, K int) [][]int {
	var ps Coor = points
	heap.Init(&ps)
	var ans [][]int
	for i := 0; i < K; i++ {
		ans = append(ans, heap.Pop(&ps).([]int))
	}
	return ans
}
