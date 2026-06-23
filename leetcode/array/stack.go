package array

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top elements of the stack.
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() T {
	var zero T // Default zero value for type T

	// Get the last element
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements[index] = zero

	// Remove the last element by slicing
	s.elements = s.elements[:index]

	return element
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() T {
	return s.elements[len(s.elements) - 1]
}

// IsEmpty checks if the stack has no elements
func(s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the current number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
