package stack

import "testing"

func TestCreateStack(t *testing.T) {
	stack := NewStack()
	if stack == nil {
		t.Fatal("something occurs.")
	}
}

func TestPushAndPop(t *testing.T) {
	stack := NewStack()
	value := "value"
	stack.Push(value)
	res := stack.Pop()
	if res != value {
		t.Error("something occurs.")
	}
}

func TestOutputOrder(t *testing.T) {
	stack := NewStack()
	val1 := "value1"
	val2 := "value2"
	stack.Push(val1)
	stack.Push(val2)
	res1 := stack.Pop()
	res2 := stack.Pop()
	if res1 != "value2" {
		t.Error("something occurs.")
	}
	if res2 != "value1" {
		t.Error("something occurs.")
	}
}
