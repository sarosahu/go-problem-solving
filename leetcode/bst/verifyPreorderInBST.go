package bst

import "math"

/**
 * Given a array of preOrder sequences of a Binary Tree.
 * Validate from the order if the tree is BST or not.
 */


/**
 * Monotonic Stack
	Intuition
	>>>>>>>>>
	In a BST preorder traversal, we visit root, then left subtree, then right subtree. 
	When we move to a right subtree, all subsequent values must be greater than the ancestors 
	we are leaving behind. The key insight is to use a decreasing stack to track ancestors. 
	When we encounter a larger value, we pop smaller ancestors and update the minimum limit, as we are now in a right subtree.

	Algorithm
	>>>>>>>>>
	1. Initialize a stack and set min_limit to negative infinity.
	2. Iterate through each number in the preorder sequence.
	3. While the stack is not empty and the stack top is less than the current number, pop from the stack and update min_limit to the popped value.
	4. If the current number is less than or equal to min_limit, return false as it violates BST property.
	5. Push the current number onto the stack.
	6. If all numbers are processed without violations, return true.
 */
func verifyPreorder(preorder []int) bool {
	minLimit := math.MinInt32
	stack := &Stack[int]{}

	for _, num := range preorder {
		for !stack.IsEmpty() {
			if topVal, _ := stack.Peek(); num > topVal {
				minLimit = topVal
				stack.Pop()
			} else {
				break
			}
		}

		if num <= minLimit {
			return false
		}
		stack.Push(num)
	}
	return true
}

/**
 * Recursion
	Intuition
	>>>>>>>>>
	We can verify the preorder sequence by simulating the construction of the BST. 
	Each recursive call attempts to build a subtree within given bounds. The key insight 
	is that for a valid preorder sequence, we can greedily consume elements that fall within 
	the current subtree's valid range, recursively processing left and right subtrees.

 * Algorithm
   >>>>>>>>>
	1. Maintain a global index to track the current position in the preorder array.
	2. Create a recursive helper function that takes min_limit and max_limit as boundaries.
	3. If the index reaches the end, return true as all elements have been processed.
	4. Check if the current element falls within the valid range. If not, return false.
	5. Increment the index and recursively verify the left subtree (with max_limit as current value) and right subtree (with min_limit as current value).
	6. Return true if either the left or right subtree verification succeeds, allowing the sequence to be valid.
 */

func verifyPreorderR(preorder []int) bool {
	i := 0
	return verifyPreorderHelper(preorder, &i, math.MinInt32, math.MaxInt32)
}

func verifyPreorderHelper(preorder []int, idx *int, minLimit int, maxLimit int) bool {
	currIdx := (*idx)
	if currIdx == len(preorder) {
		return true
	}
	currRoot := preorder[currIdx]
	if currRoot <= minLimit || currRoot >= maxLimit {
		return false
	}

	(*idx)++
	leftRes := verifyPreorderHelper(preorder, idx, minLimit, currRoot)
	rightRes := verifyPreorderHelper(preorder, idx, currRoot, maxLimit)

	return leftRes || rightRes
}