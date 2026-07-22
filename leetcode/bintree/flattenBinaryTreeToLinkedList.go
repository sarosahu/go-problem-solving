package bintree

/**
 * 114. Flatten Binary Tree to Linked List

 * Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points 
to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.

Ex 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [0]
Output: [0]
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
 */


func flatten(root *TreeNode)  {
	flattenUStack(root)
}

func flattenUStack(root *TreeNode) {
	stk := Stack[*TreeNode]{}
	curr := root
	var prev *TreeNode = nil
	for curr != nil || !stk.IsEmpty() {
		if curr != nil {
			if curr.Right != nil {
				stk.Push(curr.Right)
			}
			next := curr.Left
			curr.Left = nil
			curr.Right = next
			prev = curr
			curr = next
		} else {
			curr, _ = stk.Pop()
			prev.Right = curr
		}
	}
}

func flattenE(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			rightMost := curr.Left
			for rightMost.Right != nil {
				rightMost = rightMost.Right
			}
			rightMost.Right = curr.Right
            curr.Right = curr.Left
            curr.Left = nil
		}
		curr = curr.Right
	}
}