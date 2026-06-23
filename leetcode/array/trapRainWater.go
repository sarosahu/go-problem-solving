package array

func trapUsingMonoStack(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	idx := 0
	stack := Stack[int]{}
	area := 0
	for idx < n {
		for !stack.IsEmpty() && height[idx] > height[stack.Peek()] {
			top := stack.Pop()
			if stack.IsEmpty() {
				break
			}
			dist := idx - stack.Peek() - 1
			boundedHeight := min(height[idx], height[stack.Peek()]) - height[top]
			area += dist * boundedHeight
		}
		stack.Push(idx)
		idx++
	}
	return area
}

func trapUsingDP(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    leftMax := make([]int, n)
    rightMax := make([]int, n)

    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i-1])
    }
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i+1])
    }
    area := 0
    for i := 1; i < n - 1; i++ {
        minHeight := min(leftMax[i], rightMax[i])
        if height[i] < minHeight {
            area += minHeight - height[i]
        }
    }
    return area
}

func trap2P(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    leftMax, rightMax := 0,0
    left, right := 0, n - 1
    area := 0
    for left < right {
        leftMax = max(leftMax, height[left])
        rightMax = max(rightMax, height[right])
        if leftMax < rightMax {
            area += leftMax - height[left]
            left++
        } else {
            area += rightMax - height[right]
            right--
        }
    }
    return area
}