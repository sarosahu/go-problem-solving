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
