package heaps

import (
	"container/heap"
	"sort"
)

// 10.1 Merge Sorted Files
func MergeSortedArrays(sortedArrays [][]int) []int {

	var lists Values
	for _, list := range sortedArrays {
		if len(list) > 0 {
			heap.Push(&lists, Val{
				item:  list[0],
				array: list[1:],
			})
		}
	}

	var result []int
	for lists.Len() > 0 {
		curr := heap.Pop(&lists).(Val)
		result = append(result, curr.item)

		if len(curr.array) > 0 {
			heap.Push(&lists, Val{
				item:  curr.array[0],
				array: curr.array[1:],
			})
		}
	}
	return result
}

type Val struct {
	item  int
	array sort.IntSlice
}

type Values []Val

func (v Values) Len() int {
	return len(v)
}

func (v Values) Less(i, j int) bool {
	return v[i].item < v[j].item
}

func (v Values) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v *Values) Push(x interface{}) {
	*v = append(*v, x.(Val))
}

func (v *Values) Pop() interface{} {
	x := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return x
}

type MySlice struct {
	sort.IntSlice
}

func (m *MySlice) Pop() interface{} {
	panic("implement me")
}

func (m *MySlice) Push(x interface{}) {

}

func MyFunc() {
	myVar := &MySlice{}
	//myIntSlice := &myVar.IntSlice
	heap.Init(myVar)
}
