package stack

type Stack interface {
	Empty() bool
	Max() interface{}
	Pop() interface{}
	Push(x interface{})
	Peek() interface{}
}
