/*
128. Longest Consecutive Sequence

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
Example 3:

Input: nums = [1,0,1,2]
Output: 3
 

Constraints:

0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
 */
package hashtable

func LongestConsecutive(nums []int) int {
    return longestConsecutiveUsingSet(nums)
}

func longestConsecutiveUsingSet(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    set := make(map[int]struct{})
    for _, num := range(nums) {
        set[num] = struct{}{}
    }
    maxLen := 1
    for _, num := range(nums) {
        _, exist := set[num]
        if (!exist) {
            continue
        }
        delete(set, num)
        currLen := 1
        currNum := num
        _, exist = set[currNum - 1]
        for exist == true {
            currLen += 1
            currNum -= 1
            delete(set, currNum)
            _, exist = set[currNum - 1]
        }
        currNum = num
        _, exist = set[currNum + 1]
        for exist == true {
            currLen += 1
            currNum += 1
            delete(set, currNum)
            _, exist = set[currNum + 1]
        }
        maxLen = max(maxLen, currLen)
    }
    return maxLen
}