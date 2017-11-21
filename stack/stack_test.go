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
	res, err := stack.Pop()
	if err != nil {
		t.Error("something occurs.")
	}
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
	res1, err1 := stack.Pop()
	if err1 != nil {
		t.Error("something occurs.")
	}
	if res1 != "value2" {
		t.Error("something occurs.")
	}
	res2, err2 := stack.Pop()
	if err2 != nil {
		t.Error("something occurs.")
	}
	if res2 != "value1" {
		t.Error("something occurs.")
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := NewStack()
	res, err := stack.Pop()
	if err == nil {
		t.Error("something occurs.")
	}
	if err.Error() != "error: pop from empty list" {
		t.Error("something occurs.")
	}
	if res != "" {
		t.Error("something occurs.")
	}
}
