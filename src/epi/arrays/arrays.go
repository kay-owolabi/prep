package arrays

type Color int

const (
	RED   Color = 0
	WHITE Color = 1
	BLUE  Color = 2
)

func dutchFlagPartition(pivotIndex int, aPtr []Color) {
	pivot := aPtr[pivotIndex]
	length := len(aPtr)
	for smaller, equal, larger := 0, 0, length-1; equal < larger; {
		if aPtr[equal] < pivot {
			aPtr[equal], aPtr[smaller] = aPtr[smaller], aPtr[equal]
			smaller++
			equal++
		} else if aPtr[equal] == pivot {
			equal++
		} else {
			aPtr[equal], aPtr[larger] = aPtr[larger], aPtr[equal]
			larger--
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
