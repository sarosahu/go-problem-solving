package bintree

func PostorderTraversal(root *TreeNode) []int {
    curr := root
	stack := &Stack[*TreeNode]{}
	var prev *TreeNode
	ans := make([]int, 0)

	for curr != nil || !stack.IsEmpty() {
		if curr != nil {
			stack.Push(curr)
			curr = curr.Left
		} else {
			curr, _ = stack.Peek()
			if curr.Right == nil || curr.Right == prev {
				ans = append(ans, curr.Val)
				stack.Pop()
				prev = curr
				curr = nil
			} else {
				curr = curr.Right
			}
		}
	}
	return ans
}