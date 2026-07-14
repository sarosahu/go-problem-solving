package stack

/**
 * 84. Largest Rectangle in Histogram

 * Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, 
 * return the area of the largest rectangle in the histogram.
 */

// Brute force approach
func largestRectangleAreaBF(heights []int) int {
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		currHeight := heights[i]
		leftIter, rightIter := i, i

		// Expand to the left
		for leftIter-1 >= 0 && currHeight <= heights[leftIter-1] {
			leftIter--
		}

		// Expand to the right
		for rightIter+1 < len(heights) && currHeight <= heights[rightIter+1] {
			rightIter++
		}

		// Calculate current rectangle area
		currArea := (rightIter - leftIter + 1) * currHeight

		// Update max area
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}

func largestRectangleAreaMonoStk1(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	var stk []int

	// Nearest smaller to left
	for i := 0; i < n; i++ {
		// while !stk.isEmpty() && heights[stk.peek()] >= heights[i]
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1] // pop
		}
		
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1] // peek
		}
		
		stk = append(stk, i) // push
	}

	// Clear the stack for the next pass
	stk = stk[:0]

	// Nearest smaller to right
	for i := n - 1; i >= 0; i-- {
		// while !stk.isEmpty() && heights[stk.peek()] >= heights[i]
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1] // pop
		}
		
		if len(stk) == 0 {
			right[i] = n
		} else {
			right[i] = stk[len(stk)-1] // peek
		}
		
		stk = append(stk, i) // push
	}

	// Calculate the maximum area
	maxArea := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		currArea := heights[i] * width
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}


func largestRectangleAreaMonoStk2(heights []int) int {
	length := len(heights)
	
	// Pre-allocate stack capacity to avoid mid-loop resizing overhead
	stack := make([]int, 0, length+1)
	stack = append(stack, -1) // push -1
	
	maxArea := 0

	for i := 0; i < length; i++ {
		// while stack.peek() != -1 && heights[stack.peek()] >= heights[i]
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			// Pop current element
			currentHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			
			// Peek the next element for calculating width
			currentWidth := i - stack[len(stack)-1] - 1
			
			currArea := currentHeight * currentWidth
			if currArea > maxArea {
				maxArea = currArea
			}
		}
		stack = append(stack, i) // push i
	}

	// Process remaining elements left in the stack
	// while stack.peek() != -1
	for stack[len(stack)-1] != -1 {
		currentHeight := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1] // pop
		
		currentWidth := length - stack[len(stack)-1] - 1
		
		currArea := currentHeight * currentWidth
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}


