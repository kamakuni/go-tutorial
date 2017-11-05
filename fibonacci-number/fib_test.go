package main

import (
	"testing"
)

func TestFibWith0(t *testing.T) {
	res := Fib(0)
	if res != 0 {
		t.Error("fail")
	}
}

func TestFibWith1(t *testing.T) {
	res := Fib(1)
	if res != 1 {
		t.Error("fail")
	}
}

func TestFibWith2(t *testing.T) {
	res := Fib(2)
	if res != 1 {
		t.Error("fail")
	}
}

func TestFibWith10(t *testing.T) {
	res := Fib(10)
	if res != 55 {
		t.Error("fail")
	}
}
