package searching

import "sort"

// 11.1 Search A Sorted Array For First Occurrence Of K

func SearchFirstOfK(a []int, k int) int {
	sortedInts := sort.IntSlice(a)
	res := sortedInts.Search(k)

	if res < sortedInts.Len() && sortedInts[res] == k {
		return res
	}
	return -1
}

func BSearch(t int, A []int) int {
	L, U := 0, len(A)-1
	for L <= U {
		M := L + (U-L)/2
		if A[M] < t {
			L = M + 1
		} else if A[M] == t {
			return M
		} else {
			U = M - 1
		}
	}
	return -1
}

// 23280720657753
