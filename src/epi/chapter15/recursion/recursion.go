package recursion

import (
	"epi/test"
	"fmt"
)

func NQueens(n int) []test.IntSlice {
	var result []test.IntSlice
	solveNQueens(n, 0, &[]int{}, &result)
	return result
}

func solveNQueens(n, row int, colPlacement *[]int, result *[]test.IntSlice) {
	if row == n {
		res := make([]int, len(*colPlacement))
		copy(res, *colPlacement)
		*result = append(*result, res)
	} else {
		for col := 0; col < n; col++ {
			*colPlacement = append(*colPlacement, col)
			if isValid(*colPlacement) {
				solveNQueens(n, row+1, colPlacement, result)
			}
			*colPlacement = (*colPlacement)[:len(*colPlacement)-1]
		}
	}
}

func isValid(colPlacement []int) bool {
	rowID := len(colPlacement) - 1
	for i := 0; i < rowID; i++ {
		if diff := abs(colPlacement[i] - colPlacement[rowID]); diff == 0 || diff == rowID-i {
			return false
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// 15.1 The Towers of Hanoi problem
func ComputeTowerHanoi(numRings int) [][]int {
	var result [][]int
	var moveRings func(n, src, dst, tmp int)
	moveRings = func(n, src, dst, tmp int) {
		if n > 0 {
			moveRings(n-1, src, tmp, dst)
			result = append(result, []int{src, dst})
			moveRings(n-1, tmp, dst, src)
		}
	}
	moveRings(numRings, 0, 1, 2)

	/*
		//type funcStack struct{ n, src, dst, tmp int }
		//Stack, pop := []funcStack{{numRings, 0, 1, 2}}, false
		//for len(Stack) > 0 {
		//	curr := Stack[len(Stack)-1]
		//	if curr.n-1 > 0 && !pop {
		//		Stack = append(Stack, funcStack{curr.n - 1, curr.src, curr.tmp, curr.dst})
		//		continue
		//	}
		//	Stack, pop, result = Stack[:len(Stack)-1], true, append(result, []int{curr.src, curr.dst})
		//	if curr.n-1 > 0 {
		//		Stack, pop = append(Stack, funcStack{curr.n - 1, curr.tmp, curr.dst, curr.src}), false
		//	}
		//}
	*/
	return result
}

// 15.0 The Towers of Hanoi problem
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}

func RecPermute(soFar, rest string) {
	if rest == "" {
		fmt.Printf("out:->\t%v\n", soFar)
	} else {
		for i := 0; i < len(rest); i++ {
			next := soFar + rest[i:i+1]
			remaining := rest[0:i] + rest[i+1:]
			RecPermute(next, remaining)
		}
	}
}

func ListPermutations(s string) {
	RecPermute("", s)
}

func RecSubsets(soFar, rest string) {
	if rest == "" {
		fmt.Printf("out:->\t%s\n", soFar)
	} else {
		RecSubsets(soFar+rest[0:1], rest[1:])
		RecSubsets(soFar, rest[1:])
	}
}

func ListSubsets(str string) {
	RecSubsets("", str)
}
