package dynamicprogramming

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	fMinus2, fMinus1 := 0, 1

	for i := 1; i < n; i++ {
		f := fMinus2 + fMinus1
		fMinus2, fMinus1 = fMinus1, f
	}
	return fMinus1
}

func FindMaximumSubarray(A []int) int {
	maxSeen, maxEnd := 0, 0
	for _, a := range A {
		maxEnd = Max(a, maxEnd+a)
		maxSeen = Max(maxSeen, maxEnd)
	}
	return maxSeen
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func NumCombinationsForFinalScore(finalScore int, individualPlayScores []int) int {
	numCombinationsForScore := make([][]int, len(individualPlayScores))
	for i, _ := range numCombinationsForScore {
		numCombinationsForScore[i] = make([]int, finalScore+1)
	}

	for i := 0; i < len(individualPlayScores); i++ {
		numCombinationsForScore[i][0] = 1
		for j := 1; j <= finalScore; j++ {
			var withoutThisPlay, withThisPlay int
			if i >= 1 {
				withoutThisPlay = numCombinationsForScore[i-1][j]
			}

			if j >= individualPlayScores[i] {
				withThisPlay = numCombinationsForScore[i][j-individualPlayScores[i]]
			}

			numCombinationsForScore[i][j] = withoutThisPlay + withThisPlay
		}
	}
	return numCombinationsForScore[len(numCombinationsForScore)-1][finalScore]
}
