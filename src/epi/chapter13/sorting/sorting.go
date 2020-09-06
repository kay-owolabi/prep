package sorting

import (
	"sort"
)

// 13.1 Compute The Intersection of Two Sorted Arrays
func IntersectTwoSortedArrays(a, b []int) []int {
	result := []int{}

	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] < b[j] {
			i++
		} else if a[i] == b[j] {
			n := len(result)
			if n == 0 || result[n-1] != a[i] {
				result = append(result, a[i])
			}
			i, j = i+1, j+1
		} else {
			j++
		}
	}
	return result
}

type Student struct {
	name string
	gPA  float32
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].name < s[j].name }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Students) Sort()              { sort.Sort(s) }
