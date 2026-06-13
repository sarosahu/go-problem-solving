package bintree

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top elements of the stack.
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	var zero T // Default zero value for type T
	if s.IsEmpty() {
		return zero, errors.New("Stack is empty")
	}

	// Get the last element
	index := len(s.elements) - 1
	element := s.elements[index]

	// Remove the last element by slicing
	s.elements = s.elements[:index]

	return element, nil
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("Stack is empty")
	}
	return s.elements[len(s.elements) - 1], nil
}

// IsEmpty checks if the stack has no elements
func(s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the current number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	intStack.Push(40)
	intStack.Push(50)
	fmt.Println("Stack size", intStack.Size())

	val, _ := intStack.Pop()
	fmt.Println("Popped:", val)

	top, _ := intStack.Peek()
	fmt.Println("Top value:", top)
}