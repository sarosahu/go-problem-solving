package backtrack

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"testing"
)

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

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "Three elements (Standard LeetCode case)",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 2, 1}, {3, 1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)

			// Sort both expected and actual results to make comparison order-independent
			sortPermutations(got)
			sortPermutations(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to sort a 2D slice lexicographically for stable comparison
func sortPermutations(arr [][]int) {
	sort.Slice(arr, func(i, j int) bool {
		for k := 0; k < len(arr[i]) && k < len(arr[j]); k++ {
			if arr[i][k] != arr[j][k] {
				return arr[i][k] < arr[j][k]
			}
		}
		return len(arr[i]) < len(arr[j])
	})
}

func TestPermuteUniqueBacktrack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Empty Array",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "Single Element",
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "Array with Duplicates",
			input: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name:  "All Elements Same",
			input: []int{7, 7, 7, 7, 7, 7, 7, 7, 7},
			expected: [][]int{
				{7, 7, 7, 7, 7, 7, 7, 7, 7},
			},
		},
		{
			name:  "Distinct Elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := permuteUniqueBacktrack(tt.input)

			if len(actual) != len(tt.expected) {
				t.Fatalf("For %v: expected length %d, got %d", tt.input, len(tt.expected), len(actual))
			}

			// Validate that every generated permutation exists inside our expectation pool
			for _, perm := range actual {
				found := false
				for _, exp := range tt.expected {
					if slices.Equal(perm, exp) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("For %v: produced unexpected permutation %v", tt.input, perm)
				}
			}
		})
	}
}

func TestPermuteUsingNextPermutationLogic(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Empty Array",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "Single Element",
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "Array with Duplicates",
			input: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name:  "All Elements Same",
			input: []int{7, 7, 7, 7, 7, 7, 7, 7, 7},
			expected: [][]int{
				{7, 7, 7, 7, 7, 7, 7, 7, 7},
			},
		},
		{
			name:  "Distinct Elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := permuteUniqueUsingNextPermutation(tt.input)

			if len(actual) != len(tt.expected) {
				t.Fatalf("For %v: expected length %d, got %d", tt.input, len(tt.expected), len(actual))
			}

			// Validate that every generated permutation exists inside our expectation pool
			for _, perm := range actual {
				found := false
				for _, exp := range tt.expected {
					if slices.Equal(perm, exp) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("For %v: produced unexpected permutation %v", tt.input, perm)
				}
			}
		})
	}
}

func TestGeneratePalindromes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Impossible Palindrome (Too many odd counts)",
			input:    "abc",
			expected: []string{},
		},
		{
			name:     "Single Character",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "Simple Even Palindrome",
			input:    "aabb",
			expected: []string{"abba", "baab"},
		},
		{
			name:     "Simple Odd Palindrome",
			input:    "aabbh",
			expected: []string{"abhba", "bahab"},
		},
		{
			name:     "All Same Characters",
			input:    "aaaa",
			expected: []string{"aaaa"},
		},
		{
			name:     "Empty String",
			input:    "",
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generatePalindromes(tt.input)

			if len(actual) != len(tt.expected) {
				t.Fatalf("For string %q: expected length %d, got %d", tt.input, len(tt.expected), len(actual))
			}

			// Verify that every generated palindrome matches an expected entry
			for _, generatedItem := range actual {
				if !slices.Contains(tt.expected, generatedItem) {
					t.Errorf("For string %q: produced unexpected palindrome %q", tt.input, generatedItem)
				}
			}
		})
	}
}

// Helper function to turn a slice of slices into a standardized string map signature
// This allows order-independent comparison of the results.
func makePowerSetMap(set [][]int) map[string]bool {
	m := make(map[string]bool)
	for _, sub := range set {
		// Clone and sort the subset so [1, 2] and [2, 1] look identical
		cloned := make([]int, len(sub))
		copy(cloned, sub)
		slices.Sort(cloned)
		m[fmt.Sprintf("%v", cloned)] = true
	}
	return m
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Empty Set",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "Single Element",
			input: []int{0},
			expected: [][]int{
				{},
				{0},
			},
		},
		{
			name:  "Standard Three Element Set",
			input: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1}, {2}, {3},
				{1, 2}, {1, 3}, {2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := subsets(tt.input)

			if len(actual) != len(tt.expected) {
				t.Fatalf("For %v: expected power set size %d, got %d", tt.input, len(tt.expected), len(actual))
			}

			actualMap := makePowerSetMap(actual)
			expectedMap := makePowerSetMap(tt.expected)

			// Cross-verify all keys match
			for key := range actualMap {
				if !expectedMap[key] {
					t.Errorf("For %v: generated unexpected subset %s", tt.input, key)
				}
			}
		})
	}
}

// Helper function to turn a slice of slices into an order-independent string map
func makeCombinationMap(combinations [][]int) map[string]bool {
	m := make(map[string]bool)
	for _, comb := range combinations {
		// Clone and sort the individual combination so [2, 3] and [3, 2] look identical
		cloned := make([]int, len(comb))
		copy(cloned, comb)
		slices.Sort(cloned)
		m[fmt.Sprintf("%v", cloned)] = true
	}
	return m
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Standard Test Case",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name:       "Multiple Reuse Options",
			candidates: []int{2, 3, 5},
			target:     8,
			expected: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name:       "No Combinations Possible",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Target Matches Single Element",
			candidates: []int{1},
			target:     1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:       "Target Matches Single Element with Larger Numbers Available",
			candidates: []int{1},
			target:     2,
			expected: [][]int{
				{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := combinationSum(tt.candidates, tt.target)

			if len(actual) != len(tt.expected) {
				t.Fatalf("For candidates %v and target %d: expected length %d, got %d", 
					tt.candidates, tt.target, len(tt.expected), len(actual))
			}

			actualMap := makeCombinationMap(actual)
			expectedMap := makeCombinationMap(tt.expected)

			// Verify that every combination found matches the expected pool
			for key := range actualMap {
				if !expectedMap[key] {
					t.Errorf("For candidates %v and target %d: produced unexpected combination %s", 
						tt.candidates, tt.target, key)
				}
			}
		})
	}
}
