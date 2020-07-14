package arrays

import "math"

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
