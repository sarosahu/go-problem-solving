/*
*

  - 314. Binary Tree VerMcal Order Traversal
    Given the root of a binary tree, return the ver;cal order traversal of its nodes' values. (i.e.,
    from top to bobom, column by column).
    If two nodes are in the same row and column, the order should be from ler to right.

    Ex 1:
    Input: root = [3,9,20,null,null,15,7]
    Output: [[9],[3,15],[20],[7]]

    Ex 2:
    Input: root = [3,9,8,4,0,1,7]
    Output: [[4],[9],[3,0,1],[8],[7]]

    Ex 3:
    Input: root = [1,2,3,4,10,9,11,null,5,null,null,null,null,null,null,null,6]
    Output: [[4],[2,5],[1,10,9,6],[3],[11]]

    Constraints:

  - The number of nodes in the tree is in the range [0, 100].

  - -100 <= Node.val <= 100
*/
package bintree

import "slices"

type TreeNodeWithColumn struct{
	node *TreeNode
	column int
}
func verticalOrder(root *TreeNode) [][]int {
	output := [][]int {}
	if root == nil {
		return output
	}

	columnTable := map[int][]int{}
	queue := Queue[*TreeNodeWithColumn]{}
	currNodeInfo := TreeNodeWithColumn{
		node: root,
		column: 0,
	}
	queue.Enqueue(&currNodeInfo)
	for !queue.IsEmpty() {
		curr, _ := queue.Dequeue()
		currNode := curr.node
		column := curr.column

		//if currNode != nil {
		if _, exist := columnTable[column]; !exist {
			columnTable[column] = []int{}
		}
		columnTable[column] = append(columnTable[column], currNode.Val)
		if currNode.Left != nil {
			queue.Enqueue(&TreeNodeWithColumn{
				node: currNode.Left,
				column: column - 1,
			})
		}
		if currNode.Right != nil {
			queue.Enqueue(&TreeNodeWithColumn{
				node: currNode.Right,
				column: column + 1,
			})
		}
		//}
	}

	keys := make([]int, 0, len(columnTable))
	for k := range columnTable {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, col := range keys {
		output = append(output, columnTable[col])
	}
	return output
}

// This is efficient than above. We don't need to do sorting of keys
// We just need to track the min column index and max column index.
func verticalOrderE(root *TreeNode) [][]int {
	output := [][]int {}
	if root == nil {
		return output
	}

	columnTable := map[int][]int{}
	queue := Queue[*TreeNodeWithColumn]{}
	currNodeInfo := TreeNodeWithColumn{
		node: root,
		column: 0,
	}
	queue.Enqueue(&currNodeInfo)
	minCol, maxCol := 0,0
	for !queue.IsEmpty() {
		curr, _ := queue.Dequeue()
		currNode := curr.node
		column := curr.column
		minCol, maxCol = min(minCol, column), max(maxCol, column)

		//if currNode != nil {
		if _, exist := columnTable[column]; !exist {
			columnTable[column] = []int{}
		}
		columnTable[column] = append(columnTable[column], currNode.Val)
		if currNode.Left != nil {
			queue.Enqueue(&TreeNodeWithColumn{
				node: currNode.Left,
				column: column - 1,
			})
		}
		if currNode.Right != nil {
			queue.Enqueue(&TreeNodeWithColumn{
				node: currNode.Right,
				column: column + 1,
			})
		}
		//}
	}

	for i := minCol; i <= maxCol; i++ {
		output = append(output, columnTable[i])
	}
	return output
}
