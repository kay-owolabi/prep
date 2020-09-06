package arrays

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Color int

const (
	RED   Color = 0
	WHITE Color = 1
	BLUE  Color = 2
)

//5.1
func DutchFlagPartition(pivotIndex int, aPtr []Color) {
	pivot := aPtr[pivotIndex]
	for smaller, equal, larger := 0, 0, len(aPtr); equal < larger; {
		if aPtr[equal] < pivot {
			aPtr[equal], aPtr[smaller] = aPtr[smaller], aPtr[equal]
			smaller++
			equal++
		} else if aPtr[equal] == pivot {
			equal++
		} else {
			larger--
			aPtr[equal], aPtr[larger] = aPtr[larger], aPtr[equal]
		}
	}
}

func plusOne(slice []int) (ret []int) {
	carryOver := true
	for i := len(slice) - 1; i >= 0 && carryOver; i-- {
		if carryOver {
			slice[i]++
		}
		carryOver = slice[i] >= 10
		slice[i] %= 10
	}

	if carryOver {
		slice = append([]int{1}, slice...)
	}

	return slice
}

func multiply(nums1, nums2 []int) []int {
	panic("implement me")
}

func canReach(maxAdvanceSteps []int) bool {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	steps := 0
	for i := 0; i < len(maxAdvanceSteps)-1; i, steps = i+1, steps-1 {
		if steps < 0 {
			return false
		}

		steps = max(maxAdvanceSteps[i], steps)
	}
	return true
}

//5.6 Buy and Sell Stock Once
func ComputeMaxProfit(prices []float64) float64 {
	maxProfit, minSoFar := 0, math.MaxInt64
	for _, price := range prices {
		price := int(price * 1000)
		if minSoFar > price {
			minSoFar = price
		} else if maxProfit < price-minSoFar {
			maxProfit = price - minSoFar
		}
	}
	x := float64(maxProfit) / 1000
	//res := math.Round(x/0.05) * 0.05
	return x
}

// 5.12 Sample Offline Data
func RandomSampling(k int, A []int) {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < k; i++ {
		r := gen.Intn(len(A))
		A[i], A[r] = A[r], A[i]
	}
}

// 5.18 Compute The Spiral Ordering of a 2D array
func MatrixSpiralOrder(squareMatrix [][]int) []int {
	n := len(squareMatrix)
	SHIFT := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dir, x, y int
	var spiralOrdering []int

	for i := 0; i < (n * n); i++ {
		spiralOrdering = append(spiralOrdering, squareMatrix[x][y])
		squareMatrix[x][y] = 0
		nextX, nextY := x+SHIFT[dir][0], SHIFT[dir][1]
		if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || squareMatrix[nextX][nextY] == 0 {
			dir = (dir + 1) % 4
			nextX, nextY = x+SHIFT[dir][0], SHIFT[dir][1]
		}
		x, y = nextX, nextY
	}
	return spiralOrdering
}

func MatrixSpiralOrderOld(squareMatrix [][]int) []int {
	n := len(squareMatrix)
	var spiralOrdering []int
	matrixLayerInClockwise := func(offset int) {
		if offset == n-offset-1 {
			return
		}

		for j := offset; j < n-offset-1; j++ {
			spiralOrdering = append(spiralOrdering, squareMatrix[offset][j])
		}

		for i := offset; i < n-offset-1; i++ {
			spiralOrdering = append(spiralOrdering, squareMatrix[i][n-1-offset])
		}
		for j := n - 1 - offset; j > offset; j-- {
			spiralOrdering = append(spiralOrdering, squareMatrix[n-1-offset][j])
		}

		for i := n - 1 - offset; i > offset; i-- {
			spiralOrdering = append(spiralOrdering, squareMatrix[i][offset])
		}
	}
	for offset := 0; offset < int(math.Ceil(0.5*float64(n))); offset++ {
		matrixLayerInClockwise(offset)
	}
	return spiralOrdering
}

func SpiralOutwards(n int) [][]int {
	dir := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	x, y := 1, 0
	result := [][]int{{0, 0}, {x, y}}
	for i, pos, dis := 1, 1, 1; i <= n; i, pos = i+1, (pos+1)%4 {
		if i%2 == 0 {
			dis++
		}
		x += dir[pos][0] * dis
		y += dir[pos][1] * dis
		result = append(result, []int{x, y})
	}
	fmt.Println(result)
	return result
}
