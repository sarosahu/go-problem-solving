package graph

import (
	"reflect"
	"testing"
)

// TestFindOrder uses table-driven tests to validate various course scheduling scenarios.
func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		// Since topological sorts can have multiple valid answers,
		// we use a slice of valid options or a custom validation function.
		wantValid     [][]int 
		wantEmpty     bool
	}{
		{
			name:          "Simple linear dependency",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}}, // Course 1 depends on 0
			wantValid:     [][]int{{0, 1}},
			wantEmpty:     false,
		},
		{
			name:          "Multiple valid topological orders",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			wantValid:     [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}},
			wantEmpty:     false,
		},
		{
			name:          "Impossible schedule due to a cyclic dependency",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}}, // 1 depends on 0, 0 depends on 1
			wantValid:     nil,
			wantEmpty:     true,
		},
		{
			name:          "Single course with no prerequisites",
			numCourses:    1,
			prerequisites: [][]int{},
			wantValid:     [][]int{{0}},
			wantEmpty:     false,
		},
		{
			name:          "Multiple independent courses",
			numCourses:    3,
			prerequisites: [][]int{},
			// Any permutation of [0, 1, 2] is technically valid.
			wantValid:     [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}},
			wantEmpty:     false,
		},
		{
			name:          "Complex cycle within a subset of courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}}, // Cycle: 1 -> 2 -> 3 -> 1
			wantValid:     nil,
			wantEmpty:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrderTopologicalOrderKahn(tt.numCourses, tt.prerequisites)

			if tt.wantEmpty {
				if len(got) != 0 {
					t.Errorf("findOrder() = %v; want empty slice due to cycle", got)
				}
				return
			}

			// Verify the output matches one of the acceptable valid orders
			isValid := false
			for _, validOrder := range tt.wantValid {
				if reflect.DeepEqual(got, validOrder) {
					isValid = true
					break
				}
			}

			if !isValid {
				t.Errorf("findOrder() = %v; is not a valid topological sort. Expected one of: %v", got, tt.wantValid)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrderTopologicalOrderR(tt.numCourses, tt.prerequisites)

			if tt.wantEmpty {
				if len(got) != 0 {
					t.Errorf("findOrder() = %v; want empty slice due to cycle", got)
				}
				return
			}

			// Verify the output matches one of the acceptable valid orders
			isValid := false
			for _, validOrder := range tt.wantValid {
				if reflect.DeepEqual(got, validOrder) {
					isValid = true
					break
				}
			}

			if !isValid {
				t.Errorf("findOrder() = %v; is not a valid topological sort. Expected one of: %v", got, tt.wantValid)
			}
		})
	}
}

// TestQueue Unit tests your custom queue implementation to ensure components behave properly.
func TestQueue(t *testing.T) {
	q := Queue[int]{}

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(10)
	q.Enqueue(20)

	if q.Size() != 2 {
		t.Errorf("Expected queue size 2, got %d", q.Size())
	}

	val, ok := q.Peek()
	if !ok || val != 10 {
		t.Errorf("Peek structural failure: got %d, %t", val, ok)
	}

	val, ok = q.Dequeue()
	if !ok || val != 10 {
		t.Errorf("First dequeue failed: got %d", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 20 {
		t.Errorf("Second dequeue failed: got %d", val)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}
