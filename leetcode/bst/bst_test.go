package bst

import (
	"math"
	"testing"
)

func TestVerifyPreorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     bool
	}{
		{
			name:     "Valid BST standard case",
			preorder: []int{5, 2, 1, 3, 6},
			want:     true,
		},
		{
			name:     "Valid BST strictly increasing (Right-skewed tree)",
			preorder: []int{1, 2, 3, 4, 5},
			want:     true,
		},
		{
			name:     "Valid BST strictly decreasing (Left-skewed tree)",
			preorder: []int{5, 4, 3, 2, 1},
			want:     true,
		},
		{
			name:     "Invalid BST (Left child larger than parent)",
			preorder: []int{5, 6, 1, 3},
			want:     false,
		},
		{
			name:     "Invalid BST (Right subtree contains smaller element)",
			// 2 is in the right subtree of 3, which is invalid
			preorder: []int{5, 3, 4, 2}, 
			want:     false,
		},
		{
			name:     "Invalid BST (Deep right subtree violation)",
			// 6 is in the right subtree of 5, but 4 is placed after it
			preorder: []int{5, 2, 6, 4},
			want:     false,
		},
		{
			name:     "Single element tree",
			preorder: []int{1},
			want:     true,
		},
		{
			name:     "Empty tree",
			preorder: []int{},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := verifyPreorder(tt.preorder)
			if got != tt.want {
				t.Errorf("verifyPreorder() = %v, want %v for case %v", got, tt.want, tt.preorder)
			}
		})
	}
}

func TestVerifyPreorderR(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     bool
	}{
		{
			name:     "Valid BST standard case",
			preorder: []int{5, 2, 1, 3, 6},
			want:     true,
		},
		{
			name:     "Valid BST strictly increasing (Right-skewed tree)",
			preorder: []int{1, 2, 3, 4, 5},
			want:     true,
		},
		{
			name:     "Valid BST strictly decreasing (Left-skewed tree)",
			preorder: []int{5, 4, 3, 2, 1},
			want:     true,
		},
		{
			name:     "Invalid BST (Left child larger than parent)",
			preorder: []int{5, 6, 1, 3},
			want:     false,
		},
		{
			name:     "Invalid BST (Right subtree contains smaller element)",
			// 2 is in the right subtree of 3, which is invalid
			preorder: []int{5, 3, 4, 2}, 
			want:     false,
		},
		{
			name:     "Invalid BST (Deep right subtree violation)",
			// 6 is in the right subtree of 5, but 4 is placed after it
			preorder: []int{5, 2, 6, 4},
			want:     false,
		},
		{
			name:     "Single element tree",
			preorder: []int{1},
			want:     true,
		},
		{
			name:     "Empty tree",
			preorder: []int{},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := verifyPreorderR(tt.preorder)
			if got != tt.want {
				t.Errorf("verifyPreorder() = %v, want %v for case %v", got, tt.want, tt.preorder)
			}
		})
	}
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single Node Tree",
			root:     &TreeNode{Val: 10},
			expected: true,
		},
		{
			name: "Valid BST",
			//      5
			//     / \
			//    3   7
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
			expected: true,
		},
		{
			name: "Invalid BST - Immediate Child Violated",
			//      5
			//     / \
			//    6   7  <- 6 is not less than 5
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
			expected: false,
		},
		{
			name: "Invalid BST - Deep Ancestor Violated",
			//      5
			//     / \
			//    3   7
			//       /
			//      4    <- 4 is in right subtree of 5, but 4 < 5
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 3},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 4},
				},
			},
			expected: false,
		},
		{
			name: "Invalid BST - Duplicate Values",
			//      5
			//     / \
			//    5   7  <- Duplicate 5 violates strict less-than rule
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
			expected: false,
		},
		{
			name: "Valid BST - Integer Boundaries",
			//      0
			//     / \
			// MinInt32 MaxInt32
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: math.MinInt32},
				Right: &TreeNode{Val: math.MaxInt32},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Using your original function name
			actual := isValidBSTR(tt.root)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestIsValidBSTBfs(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single Node Tree",
			root:     &TreeNode{Val: 10},
			expected: true,
		},
		{
			name: "Valid BST",
			//      5
			//     / \
			//    3   7
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
			expected: true,
		},
		{
			name: "Invalid BST - Immediate Child Violated",
			//      5
			//     / \
			//    6   7  <- 6 is not less than 5
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
			expected: false,
		},
		{
			name: "Invalid BST - Deep Ancestor Violated",
			//      5
			//     / \
			//    3   7
			//       /
			//      4    <- 4 is in right subtree of 5, but 4 < 5
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 3},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 4},
				},
			},
			expected: false,
		},
		{
			name: "Invalid BST - Duplicate Values",
			//      5
			//     / \
			//    5   7  <- Duplicate 5 violates strict less-than rule
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
			expected: false,
		},
		{
			name: "Valid BST - Integer Boundaries",
			//      0
			//     / \
			// MinInt32 MaxInt32
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: math.MinInt32},
				Right: &TreeNode{Val: math.MaxInt32},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Using your original function name
			actual := isValidBSTBfs(tt.root)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
