package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"epi/chapter10/heaps"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func sample() {
	x := list.New()
	x.PushFront(1)
	x.PushBack(2)
	x.Remove(x.Front())
	x.Init()

	a := heaps.IntHeap{}
	heap.Init(&a)
	heap.Push(&a, 1)
	heap.Pop(&a)
	heap.Fix(&a, 2)

	w := []int{}

	sort.Slice(w, func(i, j int) bool { return w[i] < w[j] })
	sort.Search(len(w), func(i int) bool { return w[i] >= 1 })
	sort.Reverse(sort.IntSlice(w))
	//sort.Interface() -> Len() int; Less(i, j int) bool; Swap(i, j int)
	// heap.Interface() -> sort.Interface; Push(x interface{}); Pop() interface{}

	_, _ = strconv.ParseInt("", 10, 64)
	_ = math.MaxInt64
	_ = math.MinInt64
	unicode.IsSpace(34)
	//strings.

	inputFile := bufio.NewReader(os.Stdin)
	ouputFile := bufio.NewWriter(os.Stdout)
	prohibited := bufio.NewScanner(os.Stdin)

	probs := map[string]bool{}

	for word := prohibited.Text(); word != ""; word = prohibited.Text() {
		probs[strings.TrimSpace(word)] = true
	}

	var word string
	var nonWord string

	for char, _, err := inputFile.ReadRune(); err == nil; char, _, err = inputFile.ReadRune() {
		if unicode.IsLetter(char) {
			word += strconv.QuoteRuneToASCII(char)
			ouputFile.WriteString(nonWord)
			nonWord = ""
		} else {
			nonWord += strconv.QuoteRuneToASCII(char)
			if probs[word] {
				word = ""
			}
			ouputFile.WriteString(word)
			word = ""
		}
	}
	ouputFile.WriteString(word)
	ouputFile.WriteString(nonWord)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	addr := make([]string, 4)
	for i := 0; i < 4; i++ {
		addr[i] = fmt.Sprintf("%d", ip[i])
	}
	return strings.Join(addr, ".")
}

func linkedList() {
	doubleLinkedList := list.New()
	val := 0
	doubleLinkedList.PushFront(val)
	var getVal *list.Element
	getVal = doubleLinkedList.Front()
	doubleLinkedList.Remove(getVal)
}

func main0() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
	 A - 4, 19, 33, 49 l
	 B - 2, 4, 8, 10, 22, 22, 50, 90 m
	 C - 1, 2, 3, 4, 4, 26, 100, 140 n

	 O(l+m+n)

     [A[i], B[j], C[k]] - i, j, k

	Max(Abs(A[i]-B[j]), Abs(B[j] - C[k]), Abs(C[k] - A[i])) = Dist

	Dist = Max(0, 2, 2) = 2

	fine [i,j,k] that provides the least distance, i.e, Minimize Dist.

	find Item with min dis A,B -> (4,4), (19, 19)

	find min betw ((A, B).i, c), ((A,B).j, c)

	A[0] > B[0]
	{
	 A[0] - B[1]

	} A -> 19, B-> 8, C->4


	3Min(i, j, k)
	A-> 49 B->50, C-> 100
*/

type MinIndex struct {
	i, j, k int
}

func GetMinIndex(A, B, C []int) MinIndex {
	res := MinIndex{}
	minDis := math.MaxInt32

	for i, j, k := 0, 0, 0; i < len(A) && j < len(B) && k < len(C); {
		currDis, index := GetMin(A, B, C, &i, &j, &k)
		if currDis < minDis {
			res = MinIndex{i, j, k}
			minDis = currDis
		}
		*index += 1
	}
	return res
}

func GetMin(A, B, C []int, i, j, k *int) (int, *int) {
	minArr, minIndex := Min(A, B, i, j)
	_, minIndex = Min(minArr, C, minIndex, k)
	ab := Abs(A[*i] - B[*j])
	bc := Abs(B[*j] - C[*k])
	ca := Abs(C[*k] - A[*i])
	distance := IntMax(IntMax(ab, bc), ca)
	return distance, minIndex
}

func Min(A, B []int, x, y *int) ([]int, *int) {
	if A[*x] > B[*y] {
		return B, y
	}
	return A, x
}

func IntMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SubStrings(input []string, rules [][]string) error {
	rulesMap := map[string]string{}
	for _, strs := range rules { //O(r)
		rule, to := strs[0], strs[1]
		rulesMap[rule] = to
	}

	for from, _ := range rulesMap { //0(r)
		for to := rulesMap[from]; to != ""; to = rulesMap[to] { //O(r-1)
			if from == to {
				return errors.New("invalid")
			}
			rulesMap[from] = to
		}
	}

	for i, str := range input {
		input[i] = rulesMap[str] //O(i)
	}
	return nil
}

type Slice []int

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Slice) Less(i, j int) bool { return abs(s[i]) < abs(s[j]) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solution(A []int) int {
	// write your code in Go 1.4
	n := len(A)
	var forward []int
	var backward []int

	Abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Sort(Slice(A))

	for i, j, fsum, bsum := 0, n-1, 0, 0; i < n; i, j = i+1, j-1 {
		fsum += Abs(A[i])
		bsum += Abs(A[j])
		forward = append(forward, fsum)
		backward = append([]int{}, bsum)
	}

	var minIndex int
	for i, min := 0, math.MaxInt32; i < n; i++ {
		if Abs(forward[i]-backward[i]) < min {
			min, minIndex = Abs(forward[i]-backward[i]), i
		}
	}

	var res int
	for i, val := range A {
		if i <= minIndex {
			res -= Abs(val)
		} else {
			res += Abs(val)
		}
	}

	return res
}
