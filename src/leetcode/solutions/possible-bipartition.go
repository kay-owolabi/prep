package solutions

import (
	"fmt"
	"leetcode/graphs"
)

//5
//[[1,2],[3,4],[4,5],[3,5]]

func PossibleBipartition(N int, dislikes [][]int) bool {
	var x, y, haterX, haterY int
	index := make([]int, N+1)
	haters := make([]int, N+1)
	slice := make([]int, N+1)

	for i := range slice {
		slice[i] = i
		index[i] = i
	}

	if len(dislikes) < 1 {
		return true
	}

	for _, person := range dislikes {
		x = person[0]
		y = person[1]

		xRoot := graphs.Find(slice, x)
		yRoot := graphs.Find(slice, y)

		if xRoot == yRoot {
			return false
		}

		if haters[x] == 0 {
			haters[x] = y
		}

		if haters[y] == 0 {
			haters[y] = x
		}

		haterX = haters[x]
		haterY = haters[y]
		graphs.Union(slice, haterX, y)
		graphs.Union(slice, haterY, x)

		fmt.Printf("----{%d, %d}---- \n", x, y)
		fmt.Printf("%v\n", index)
		println()
		fmt.Printf("%v\n", haters)
		fmt.Printf("%v\n", slice)
		println()
		println()

	}
	return true
}
