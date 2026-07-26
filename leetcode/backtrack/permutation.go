/**
 * Given an array nums of distinct integers, return all the possible permutations.
 You can return the answer in any order.

 Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]
 

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
 */
package backtrack

func permute(nums []int) [][]int {
    res := [][]int{}
    pArr := make([]int, len(nums))
    copy(pArr, nums)
    helper(&res, pArr, 0)
    return res
}

func helper(res *[][]int, pArr []int, idx int) {
    if idx == len(pArr) - 1 {
        dArr := make([]int, len(pArr))
        copy(dArr, pArr)
        *res = append(*res, dArr)
        return
    }
    for i := idx; i < len(pArr); i++ {
        swap(pArr, i, idx)
        helper(res, pArr, idx + 1)
        swap(pArr, i, idx)
    }
}

func swap(arr []int, i, j int) {
    if i == j {
        return
    }
    arr[i], arr[j] = arr[j], arr[i]
}