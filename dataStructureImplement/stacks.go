package main

import "fmt"

// Stack is a generic stack data structure
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the last item
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	// Stack of integers
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop()) // 2

	// Stack of strings
	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop()) // world
}

// *******************************************Explanation:
// Stack[T any] defines a stack where T can be any type.
// Methods like Push and Pop operate on T and can work for any type (integers, strings, etc.).
// var zero T initializes a zero value of the generic type T.
