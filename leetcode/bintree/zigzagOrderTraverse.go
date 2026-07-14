package bintree

import "container/list"

/**
 * 103. Binary Tree Zigzag Level Order Traversal

 * Given the root of a binary tree, return the zigzag level order traversal of
 * its nodes' values. (i.e., from left to right, then right to left for the next
 * level and alternate between).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

Example 2:

Input: root = [1]
Output: [[1]]

Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
 **/

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    deque := list.New()
    deque.PushBack(root)
    var leftOrder bool = true
    var ans [][]int
    for deque.Len() != 0 {
        n := deque.Len()
        level := make([]int, 0, n)
        for i := 0; i < n; i++ {
            //var curr *TreeNode = nil
            var currElement *list.Element = nil
            if leftOrder {
                currElement = deque.Back()
            } else {
                currElement = deque.Front()
            }
            curr := currElement.Value.(*TreeNode)
            level = append(level, curr.Val)
            deque.Remove(currElement)
            if leftOrder {
                if curr.Left != nil {
                    deque.PushFront(curr.Left)
                }
                if curr.Right != nil {
                    deque.PushFront(curr.Right)
                }
            } else {
                if curr.Right != nil {
                    deque.PushBack(curr.Right)
                }
                if curr.Left != nil {
                    deque.PushBack(curr.Left)
                }
            }
        }
        leftOrder = !leftOrder
        ans = append(ans, level)
    }
    return ans
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    var result [][]int
    queue := []*TreeNode{root}
    leftToRight := true

    for len(queue) > 0 {
        levelSize := len(queue)
        levelNodes := make([]int, levelSize)

        for i := 0; i < levelSize; i++ {
            // Pop from the front of the queue
            node := queue[0]
            queue = queue[1:]

            // Calculate the correct insertion index based on direction
            var index int
            if leftToRight {
                index = i
            } else {
                index = levelSize - 1 - i
            }
            levelNodes[index] = node.Val

            // Enqueue children in standard left-to-right order
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // Add the current level to results and flip direction
        result = append(result, levelNodes)
        leftToRight = !leftToRight
    }

    return result
}
