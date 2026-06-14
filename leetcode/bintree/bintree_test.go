package bintree

import (
	"reflect"
	"testing"
)

func TestInorderTraversalR(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: []int{},
		},
		{
			name: "Single Node",
			root: &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Standard Binary Tree",
			//      1
			//       \
			//        2
			//       /
			//      3
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "Balanced Binary Tree",
			//      4
			//    /   \
			//   2     6
			//  / \   / \
			// 1   3 5   7
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := InorderTraversalR(tt.root)
			
			// reflect.DeepEqual is used to compare slices in Go
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("InorderTraversal() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestInorderTraversalStk(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: []int{},
		},
		{
			name: "Single Node",
			root: &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Standard Binary Tree",
			//      1
			//       \
			//        2
			//       /
			//      3
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			name: "Balanced Binary Tree",
			//      4
			//    /   \
			//   2     6
			//  / \   / \
			// 1   3 5   7
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := InorderTraversalStk(tt.root)
			
			// reflect.DeepEqual is used to compare slices in Go
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("InorderTraversal() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

// TestPreorderTraversal checks various binary tree shapes using table-driven tests.
func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Empty tree",
			root: nil,
			want: []int{},
		},
		{
			name: "Single node tree",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "Left skewed tree (linear chain)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 2, 3}, // Root -> Left -> Left
		},
		{
			name: "Right skewed tree (linear chain)",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 2, 3}, // Root -> Right -> Right
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			want: []int{1, 2, 4, 5, 3, 6, 7}, // Root -> Left Subtree -> Right Subtree
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestPostorderTraversal checks various binary tree shapes using table-driven tests.
func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Empty tree",
			root: nil,
			want: []int{},
		},
		{
			name: "Single node tree",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
		{
			name: "Left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: []int{3, 2, 1}, // Left -> Parent -> Root
		},
		{
			name: "Right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: []int{3, 2, 1}, // Right -> Parent -> Root
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			want: []int{4, 5, 2, 6, 7, 3, 1}, // Left Subtree -> Right Subtree -> Root
		},
		{
			name: "Asymmetric tree with missing nodes",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 5},
				},
			},
			want: []int{4, 2, 5, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostorderTraversal(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
