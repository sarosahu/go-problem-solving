package graph

import (
	"fmt"
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

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name         string
		numCourses   int
		prerequisites [][]int
		want         bool
	}{
		{
			name:         "Simple Valid Path",
			numCourses:   2,
			prerequisites: [][]int{{1, 0}}, // To take course 1, you must finish 0 first.
			want:         true,
		},
		{
			name:         "Simple Direct Cycle",
			numCourses:   2,
			prerequisites: [][]int{{1, 0}, {0, 1}}, // 0 depends on 1, and 1 depends on 0.
			want:         false,
		},
		{
			name:         "No Prerequisites (Independent Courses)",
			numCourses:   3,
			prerequisites: [][]int{},
			want:         true,
		},
		{
			name:         "Long Valid Chain",
			numCourses:   4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}}, // 0 -> 1 -> 2 -> 3
			want:         true,
		},
		{
			name:         "Indirect Long Cycle",
			numCourses:   4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}}, // 0 -> 1 -> 2 -> 3 -> 0
			want:         false,
		},
		{
			name:         "Disconnected Valid Subgraphs",
			numCourses:   5,
			prerequisites: [][]int{{1, 0}, {3, 2}}, // Two distinct independent pairs, plus course 4 completely alone
			want:         true,
		},
		{
			name:         "Self Cycle Loop",
			numCourses:   1,
			prerequisites: [][]int{{0, 0}}, // Course depends on itself
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish() = %v, want %v for prerequisites %v", got, tt.want, tt.prerequisites)
			}
		})
	}
}

func TestValidTree(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{
		{
			name:     "Standard Valid Tree",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "Graph with Cycle",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
		{
			name:     "Disconnected Components (Correct Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}, {0, 1}}, // Contains a duplicate edge / multi-graph cycle
			expected: false,
		},
		{
			name:     "Disconnected Forest (Wrong Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}},
			expected: false,
		},
		{
			name:     "Single Node (Trivial Tree)",
			n:        1,
			edges:    [][]int{},
			expected: true,
		},
		{
			name:     "Two Disconnected Nodes",
			n:        2,
			edges:    [][]int{},
			expected: false,
		},
		{
			name:     "Linear Path (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: true,
		},
		{
			name:     "Star Graph (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			expected: true,
		},
	}
	fmt.Print("Testing validTree() --> \n")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTree(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("validTree(%d, %v) = %v; want %v", tt.n, tt.edges, actual, tt.expected)
			}
		})
	}
	fmt.Print("<--DONE\n")
}

func TestValidTreeUsingStack(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{
		{
			name:     "Standard Valid Tree",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "Graph with Cycle",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
		{
			name:     "Disconnected Components (Correct Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}, {0, 1}}, // Contains a duplicate edge / multi-graph cycle
			expected: false,
		},
		{
			name:     "Disconnected Forest (Wrong Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}},
			expected: false,
		},
		{
			name:     "Single Node (Trivial Tree)",
			n:        1,
			edges:    [][]int{},
			expected: true,
		},
		{
			name:     "Two Disconnected Nodes",
			n:        2,
			edges:    [][]int{},
			expected: false,
		},
		{
			name:     "Linear Path (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: true,
		},
		{
			name:     "Star Graph (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			expected: true,
		},
	}

	fmt.Print("Testing validTreeUsingStack() --> \n")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTreeUsingStack(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("validTree(%d, %v) = %v; want %v", tt.n, tt.edges, actual, tt.expected)
			}
		})
	}
	fmt.Print("<--DONE\n")
}

func TestValidTreeBfs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{
		{
			name:     "Standard Valid Tree",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "Graph with Cycle",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
		{
			name:     "Disconnected Components (Correct Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}, {0, 1}}, // Contains a duplicate edge / multi-graph cycle
			expected: false,
		},
		{
			name:     "Disconnected Forest (Wrong Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}},
			expected: false,
		},
		{
			name:     "Single Node (Trivial Tree)",
			n:        1,
			edges:    [][]int{},
			expected: true,
		},
		{
			name:     "Two Disconnected Nodes",
			n:        2,
			edges:    [][]int{},
			expected: false,
		},
		{
			name:     "Linear Path (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: true,
		},
		{
			name:     "Star Graph (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			expected: true,
		},
	}

	fmt.Print("Testing validTreeUsingStack() --> \n")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTreeBfs(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("validTree(%d, %v) = %v; want %v", tt.n, tt.edges, actual, tt.expected)
			}
		})
	}
	fmt.Print("<--DONE\n")
}

func TestValidTreeUsingDisjointSet(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{
		{
			name:     "Standard Valid Tree",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "Graph with Cycle",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
		{
			name:     "Disconnected Components (Correct Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}, {0, 1}}, // Contains a duplicate edge / multi-graph cycle
			expected: false,
		},
		{
			name:     "Disconnected Forest (Wrong Edge Count)",
			n:        4,
			edges:    [][]int{{0, 1}, {2, 3}},
			expected: false,
		},
		{
			name:     "Single Node (Trivial Tree)",
			n:        1,
			edges:    [][]int{},
			expected: true,
		},
		{
			name:     "Two Disconnected Nodes",
			n:        2,
			edges:    [][]int{},
			expected: false,
		},
		{
			name:     "Linear Path (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			expected: true,
		},
		{
			name:     "Star Graph (Valid Tree)",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}},
			expected: true,
		},
	}

	fmt.Print("Testing validTreeUsingStack() --> \n")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTreeDS(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("validTree(%d, %v) = %v; want %v", tt.n, tt.edges, actual, tt.expected)
			}
		})
	}
	fmt.Print("<--DONE\n")
}
