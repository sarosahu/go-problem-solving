package bintree
/**
 * 102. Binary Tree Level Order Traversal

 Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

 Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
 **/

// BFS approach
func levelOrder(root *TreeNode) [][]int {
    levels := [][]int{}
    if root == nil {
        return levels
    }
	queue := Queue[*TreeNode]{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		len := queue.Size()
		level := make([]int, len)
        i := 0
		for len > 0 {
			curr, _ := queue.Dequeue()
            level[i] = curr.Val
			len--
            i++
			if curr.Left != nil {
				queue.Enqueue(curr.Left)
			}
			if curr.Right != nil {
				queue.Enqueue(curr.Right)
			}
		}
		levels = append(levels, level)
	}
    return levels
}

// DFS approach
func levelOrderDfs(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    levels := [][]int{}
    helper(root, &levels, 0)
    return levels
}

func helper(curr *TreeNode, levels *[][]int, levelIdx int) {
    if len(*levels) == levelIdx {
        *levels = append(*levels, []int{})
    }
    //level := levels[levelIdx]
    (*levels)[levelIdx] = append((*levels)[levelIdx], curr.Val)

    if curr.Left != nil {
        helper(curr.Left, levels, levelIdx + 1)
    }
    if curr.Right != nil {
        helper(curr.Right, levels, levelIdx + 1)
    }
}