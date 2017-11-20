package stack

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
func (s *stack) Pop() string {
	return s.values[len(s.values)-1]
}

/*
　Push is for pushing value into a stack
*/
func (s *stack) Push(str string) {
	s.values = append(s.values, str)
}
