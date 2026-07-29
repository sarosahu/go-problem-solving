package array

import (
	"slices"
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
    //return getPermutationUsingNextPermute(n, k)
    return getPermutationUsingFactorial(n, k)
}

func getPermutationUsingNextPermute(n int, k int) string {
    nums := []int{}
    for i := 1; i <= n; i++ {
        nums = append(nums, i)
    }
    cnt := 1
    for {
        if cnt == k {
            break;
        }
        if !nextPermute(nums) {
            return ""
        }
        cnt++
    }
    res := 0
    for _, num := range nums {
        res = res * 10 + num
    }
    return strconv.Itoa(res)
}

func nextPermute(nums []int) bool {
    n := len(nums)
    k := n - 2

    for k >= 0 && nums[k] >= nums[k + 1] {
        k--
    }

    // Array is completely descending
    if k <= -1 {
        //slices.Reverse(*nums)
        return false
    }

    for i := n - 1; i > k; i-- {
        if nums[i] > nums[k] {
            nums[i], nums[k] = nums[k], nums[i]
            break;
        }
    }
    slices.Reverse(nums[k + 1 : n])
    return true
}

// This one is more efficient, I copied from gemini.
// Explanation:
/**
 * 
 */
func getPermutationUsingFactorial(n int, k int) string {
	// Pre-calculate factorials up to n
	factorials := make([]int, n)
	factorials[0] = 1
	for i := 1; i < n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	// Create a list of available digits: [1, 2, 3, ..., n]
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// Convert k to 0-indexed for clean math division
	k--

	var sb strings.Builder

	// Determine digits one-by-one from left to right
	for i := n - 1; i >= 0; i-- {
		// Calculate how many blocks of size i! fit into k
		idx := k / factorials[i]
		k = k % factorials[i]

		// Append the matching digit
		sb.WriteString(strconv.Itoa(nums[idx]))

		// Remove the consumed digit from our available pool
		nums = append(nums[:idx], nums[idx+1:]...)
	}

	return sb.String()
}