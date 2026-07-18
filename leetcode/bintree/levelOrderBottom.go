package bintree

/**
 * 107. Binary Tree Level Order Traversal II

 * Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. 
 (i.e., from left to right, level by level from leaf to root).
 *
 * Ex 1:
 Input: root = [3,9,20,null,null,15,7]
 Output: [[15,7],[9,20],[3]]

Example 2:

Input: root = [1]
Output: [[1]]

Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

 */
func levelOrderBottom(root *TreeNode) [][]int {
    //return levelOrderBottomDfs(root)
    return levelOrderBottomBfs(root)
}

func levelOrderBottomDfs(root *TreeNode) [][]int {
    levels := [][]int{}
    if root == nil {
        return levels
    }
    levelOrderBottomDfsHelper(root, 0, &levels)
    reverse(levels)
    return levels
}

func levelOrderBottomDfsHelper(node *TreeNode, level int, levels *[][]int) {
    if len(*levels) - 1 < level {
        *levels = append(*levels, []int{})
    }
    (*levels)[level] = append((*levels)[level], node.Val)
    if node.Left != nil {
        levelOrderBottomDfsHelper(node.Left, level + 1, levels)
    }
    if node.Right != nil {
        levelOrderBottomDfsHelper(node.Right, level + 1, levels)
    }
}

func levelOrderBottomBfs(root *TreeNode) [][]int {
    levels := [][]int{}
    if root == nil {
        return levels
    }
    queue := Queue[*TreeNode]{items: []*TreeNode{root}}
    for !queue.IsEmpty() {
        currLevelSize := queue.Size()
        currLevel := make([]int, currLevelSize)
        for i := 0; i < currLevelSize; i++ {
            curr, _ := queue.Dequeue()
            currLevel[i] = curr.Val
            if curr.Left != nil {
                queue.Enqueue(curr.Left)
            }
            if curr.Right != nil {
                queue.Enqueue(curr.Right)
            }
        }
        levels = append(levels, currLevel)
    }
    reverse(levels)
    return levels
}

func reverse(arr [][]int) {
    for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}