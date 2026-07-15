package backtrack

import "testing"

func TestUniquePathsIII(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Case 1: Standard Grid (Example 1)",
			grid: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 2, -1},
			},
			expected: 2,
		},
		{
			name: "Case 2: Grid with Obstacles (Example 2)",
			grid: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			},
			expected: 4,
		},
		{
			name: "Case 3: No Valid Path (Example 3)",
			grid: [][]int{
				{0, 1},
				{2, 0},
			},
			expected: 0, // Cannot walk over all non-obstacle squares exactly once
		},
		{
			name: "Case 4: Smallest Possible Grid",
			grid: [][]int{
				{1, 2},
			},
			expected: 1, // Only 1 step needed, start to end directly
		},
		{
			name: "Case 5: Linear Grid with Obstacle",
			grid: [][]int{
				{1, 0, -1, 0, 2},
			},
			expected: 0, // Blocked by the obstacle in the middle
		},
		{
			name: "Case 6: Snake/Spiral Path Requirement",
			grid: [][]int{
				{1,  0,  0},
				{-1, -1, 0},
				{2,  0,  0},
			},
			expected: 1, // Must follow a specific C-shaped path to hit all 0s
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Deep copy the grid because backtracking modifies it in place
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			actual := uniquePathsIII(gridCopy)
			if actual != tt.expected {
				t.Errorf("got %d, want %d", actual, tt.expected)
			}
		})
	}
}
