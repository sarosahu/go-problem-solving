/*
*
  - 78. Subsets
  - Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]

Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/
package backtrack

func subsets(nums []int) [][]int {
	res := [][]int{}
	comb := []int{}
	
	// Start backtracking from index 0
	backtrackSubsets(0, nums, comb, &res)
	return res
}

func backtrackSubsets(start int, nums []int, comb []int, res *[][]int) {
	// Every combination generated along the way is a valid subset!
	dst := make([]int, len(comb))
	copy(dst, comb)
	*res = append(*res, dst)

	for i := start; i < len(nums); i++ {
		// Take the element
		comb = append(comb, nums[i])
		
		// Move forward (i+1 ensures we don't reuse the same element)
		backtrackSubsets(i+1, nums, comb, res)
		
		// Backtrack (remove the element)
		comb = comb[:len(comb)-1]
	}
}

// TODO
func subsetsUsingBitmask(nums []int) [][]int {
	return nil
}