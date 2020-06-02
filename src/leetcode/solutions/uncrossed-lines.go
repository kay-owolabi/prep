package solutions

import "leetcode/lib"

/*
We write the integers of A and B (in the order they are given) on two separate horizontal lines.

Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:

A[i] == B[j];
The line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.

Return the maximum number of connecting lines we can draw in this way.
*/

func MaxUncrossedLines(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	res := make([][]int, len(A))
	for i, _ := range A {
		res[i] = make([]int, len(B))
	}

	getItem := func(i, j int) (ret int) {
		if i < 0 || j < 0 || i >= len(A) || j >= len(B) {
			return
		}
		return res[i][j]
	}

	for i := 0; i < len(A); i++ {
		a := A[i]
		for j := 0; j < len(B); j++ {
			b := B[j]
			if a == b {
				res[i][j] = 1 + getItem(i-1, j-1)
			} else {
				res[i][j] = lib.MaxInt(getItem(i-1, j), getItem(i, j-1))
			}
		}
	}

	return res[len(A)-1][len(B)-1]
}
