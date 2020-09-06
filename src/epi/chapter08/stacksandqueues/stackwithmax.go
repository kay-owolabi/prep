package stacksandqueues

type Element struct {
	item int
	max  int
}

type IntStack []Element

func (s IntStack) Max() interface{} {
	return s[len(s)-1].max
}

func (s *IntStack) Pop() interface{} {
	ret := s.Peek()
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *IntStack) Push(x interface{}) {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	currMax := x
	if !s.Empty() {
		currMax = max((*s)[len(*s)-1].max, currMax.(int))
	}

	elem := Element{
		item: x.(int),
		max:  currMax.(int),
	}

	*s = append(*s, elem)
}

func (s IntStack) Peek() interface{} {
	return s[len(s)-1].item
}

func (s IntStack) Empty() bool {
	return len(s) == 0
}
