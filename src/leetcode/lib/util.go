package lib

import "fmt"

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Print2DArray(array [][]int) {
	fmt.Println(array)
	/*for index, item := range array {
		fmt.Printf("%d : %v\n", index, item)
	}*/
}
