package bintree

/*
 * 94. Binary Tree Inorder Traversal

 * Given the root of a binary tree, return the inorder traversal of its nodes' values.

 * Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

 *Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,2,6,5,7,1,3,9,8]

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/
func InorderTraversalR(root *TreeNode) []int {
    ans := make([]int, 0)
	if root == nil {
		return ans
	}
	inorderTraversalHelper(&ans, root)

	return ans
}

func inorderTraversalHelper(ans *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	inorderTraversalHelper(ans, node.Left)
	*ans = append(*ans, node.Val)
	inorderTraversalHelper(ans, node.Right)
}

func InorderTraversalStk(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := Stack[*TreeNode]{}

	curr := root
	for curr != nil || !stack.IsEmpty() {
		if curr != nil {
			stack.Push(curr)
			curr = curr.Left
		} else {
			curr, _ = stack.Peek()
			stack.Pop()
			ans = append(ans, curr.Val)
			curr = curr.Right
		}
	}
	return ans
}