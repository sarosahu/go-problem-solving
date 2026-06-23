package bst

import "math"

/**
 * 98. Validate Binary Search Tree
 * Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys strictly less than the node's key.
The right subtree of a node contains only nodes with keys strictly greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Note: I have practiced it on 23rd June, 2026 in Java
TODO:
*/

// { Recursive approach
func isValidBST(root *TreeNode) bool {
    return isValidBSTR(root)
}

func isValidBSTR(root *TreeNode) bool {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return true;
    }

    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, minVal, maxVal int64) bool {
    if root == nil {
        return true
    }
    currVal := int64(root.Val)
    if currVal <= minVal || currVal >= maxVal {
        return false
    }
    return validate(root.Left, minVal, currVal) &&
            validate(root.Right, currVal, maxVal)
}
// }
// BFS way
type NodeWithRange struct {
    root *TreeNode
    minVal int64
    maxVal int64
}
func isValidBSTBfs(root *TreeNode) bool {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return true;
    }

    queue := Queue[NodeWithRange]{}
    queue.Enqueue(NodeWithRange{root, math.MinInt64, math.MaxInt64})

    for !queue.IsEmpty() {
        curr, _ := queue.Dequeue()
        currVal := int64(curr.root.Val)
        if currVal <= curr.minVal || currVal >= curr.maxVal {
            return false
        }
        if curr.root.Left != nil {
            queue.Enqueue(NodeWithRange{curr.root.Left, curr.minVal, currVal})
        }
        if curr.root.Right != nil {
            queue.Enqueue(NodeWithRange{curr.root.Right, currVal, curr.maxVal})
        }
    }
    return true
}