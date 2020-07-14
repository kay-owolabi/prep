package stacksandqueues

import (
	"encoding/json"
	"epi/stacksandqueues/stack"
	"epi/test"
	"fmt"
	"strconv"
	"testing"
)

type Tuple struct {
	op  string
	arg int
}

func (t *Tuple) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&t.op, &t.arg}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if got, exp := len(tmp), wantLen; got != exp {
		return fmt.Errorf("wrong number of fields in Tuple: got %d, expected %d", got, exp)
	}
	return nil
}

func TestStack_Empty(t *testing.T) {
	testData := test.GetTestData("stack_with_max.tsv")[1:]
	type args []Tuple
	type test struct {
		name string
		args args
	}
	var tests []test

	for i, datum := range testData {
		var ops args
		err := json.Unmarshal([]byte(datum[0]), &ops)
		if err != nil {
			panic(err)
		}
		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: ops,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}*/
			var s stack.Stack
			var result int
			for _, op := range tt.args {
				switch op.op {
				case "Stack":
					s = new(IntStack)
					break
				case "push":
					s.Push(op.arg)
					break
				case "pop":
					result = s.Pop().(int)
					if result != op.arg {
						t.Errorf("StackWithMax\nFailure info\n\tPop: expected %v, got %v\n", op.arg, result)
					}
					break
				case "max":
					result = s.Max().(int)
					if result != op.arg {
						t.Errorf("StackWithMax\nFailure info\n\tMax: expected %v, got %v\n", op.arg, result)
					}
					break
				case "empty":
					result = 0
					if s.Empty() {
						result = 1
					}
					if result != op.arg {
						t.Errorf("StackWithMax\nFailure info\n\tEmpty: expected %v, got %v\n", op.arg, result)
					}
					break
				default:
					t.Errorf("StackWithMax\nFailure info\n\tUnsupported stack operation: %v\n", op.op)
				}
			}
		})
	}
}
