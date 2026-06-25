package bintree

// Queue holds values of any type T
type Queue[T any] struct {
	items []T
}

// Enqueue adds an item to the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T 
		return zero, false
	}

	item := q.items[0]

	// Memory optimization: clear the reference to prevent leaks
	var zero T
	q.items[0] = zero

	q.items = q.items[1:]
	return item, true
}

// Peek returns the front item without removing it
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty returns true if the queue contains no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the current number of elements
func (q *Queue[T]) Size() int {
	return len(q.items)
}
