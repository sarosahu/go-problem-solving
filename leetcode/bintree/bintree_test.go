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

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Empty Tree (Edge Case Fix)",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "Single Node Tree",
			root: &TreeNode{Val: 1},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Standard Binary Tree",
			//      3
			//     / \
			//    9   20
			//       /  \
			//      15   7
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "Asymmetric Unbalanced Tree",
			//      1
			//     /
			//    2
			//   /
			//  3
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := levelOrder(tt.root)
			
			// DeepEqual checks multi-dimensional slices properly
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := levelOrderDfs(tt.root)
			
			// DeepEqual checks multi-dimensional slices properly
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestZigzagLevelOrderUsingList(t *testing.T) {

	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Case 1: Empty Tree (Nil Root)",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Case 2: Single Node Tree",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name: "Case 3: Unbalanced Left-Skewed Tree (Line)",
			// 1 -> 2 -> 3 -> 4
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 4: Unbalanced Right-Skewed Tree (Line)",
			// 1 -> 2 -> 3 -> 4
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 5: Perfect Symmetric Full Binary Tree",
			//      1
			//    /   \
			//   2     3
			//  / \   / \
			// 4   5 6   7
			root: &TreeNode{
				Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			expected: [][]int{
				{1},
				{3, 2},
				{4, 5, 6, 7},
			},
		},
		{
			name: "Case 6: Deep Asymmetric Tree with Missing Children",
			//        1
			//       /
			//      2
			//       \
			//        3
			//       /
			//      4
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 7: Tree with Negative and Zero Values",
			//      0
			//    /   \
			//  -5     10
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -5},
				Right: &TreeNode{Val: 10},
			},
			expected: [][]int{
				{0},
				{10, -5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_LinkedList", func(t *testing.T) {
			actual := zigzagLevelOrderUsingList(tt.root)
			if !reflect.DeepEqual(actual, tt.expected) && !(len(actual) == 0 && len(tt.expected) == 0) {
				t.Errorf("got %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestZigzagLevelOrderUsingSlices(t *testing.T) {

	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Case 1: Empty Tree (Nil Root)",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Case 2: Single Node Tree",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name: "Case 3: Unbalanced Left-Skewed Tree (Line)",
			// 1 -> 2 -> 3 -> 4
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 4: Unbalanced Right-Skewed Tree (Line)",
			// 1 -> 2 -> 3 -> 4
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 5: Perfect Symmetric Full Binary Tree",
			//      1
			//    /   \
			//   2     3
			//  / \   / \
			// 4   5 6   7
			root: &TreeNode{
				Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
			expected: [][]int{
				{1},
				{3, 2},
				{4, 5, 6, 7},
			},
		},
		{
			name: "Case 6: Deep Asymmetric Tree with Missing Children",
			//        1
			//       /
			//      2
			//       \
			//        3
			//       /
			//      4
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "Case 7: Tree with Negative and Zero Values",
			//      0
			//    /   \
			//  -5     10
			root: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -5},
				Right: &TreeNode{Val: 10},
			},
			expected: [][]int{
				{0},
				{10, -5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Slice", func(t *testing.T) {
			actual := zigzagLevelOrderUsingSlices2(tt.root)
			if !reflect.DeepEqual(actual, tt.expected) && !(len(actual) == 0 && len(tt.expected) == 0) {
				t.Errorf("got %v, want %v", actual, tt.expected)
			}
		})
	}
}

