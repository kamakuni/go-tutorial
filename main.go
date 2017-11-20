package main

import "fmt"

func main() {
	var stack = make([]string, 0)
	stack = append(stack, "test1")
	fmt.Println(len(stack))
	stack = append(stack, "test2")
	fmt.Println(len(stack))
	var new = make([]string, 0)
	copy(stack, new)
	fmt.Println(new)
	value := new[len(new)-1]
	fmt.Println(value)
}
