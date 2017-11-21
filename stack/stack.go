package stack

import "errors"

/*
 * Stack is stack
 */
type stack struct {
	values []string
}

/*
 NewStack is Constructor for a stack
*/
func NewStack() *stack {
	return &stack{values: make([]string, 0)}
}

/*
　Pop is for popping value from a stack
*/
func (s *stack) Pop() (string, error) {
	if len(s.values) == 0 {
		return "", errors.New("error: pop from empty list")
	}
	val := s.values[len(s.values)-1]
	if len(s.values) > 0 {
		s.values = s.values[:len(s.values)-1]
	}
	return val, nil
}

/*
　Push is for pushing value into a stack
*/
func (s *stack) Push(str string) {
	s.values = append(s.values, str)
}
