package bintree

func PreorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    stack := &Stack[*TreeNode]{}
    stack.Push(root)
    for !stack.IsEmpty() {
        curr, _ := stack.Pop()
        ans = append(ans, curr.Val)
        if curr.Right != nil {
            stack.Push(curr.Right)
        }
        if curr.Left != nil {
            stack.Push(curr.Left)
        }
    }
    return ans
}