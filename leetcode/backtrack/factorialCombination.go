/*
*
  - Find all possible ways to express a given integer n as a product of its factors.

Given an integer n, you need to return all possible combinations where n can be
written as a product of two or more factors. Each factor must be in the range [2, n-1],
meaning factors cannot be 1 or n itself.

For example, if n = 8:

8 can be expressed as 2 × 2 × 2
8 can also be expressed as 2 × 4
Both of these factorizations would be valid answers.
*/
package backtrack

func factorialCombination(n int) [][]int {
	res := [][]int{}
	cseq := []int{}
	// farr := [2]int{1, n}
	// res = append(res, farr[:])
	factorialCombinationHelper(&res, cseq, n, n, 2)

	return res
}

func factorialCombinationHelper(res *[][]int, cseq []int, n, remaining, start int) {
	if remaining >= start && remaining < n {
		dst := make([]int, len(cseq) + 1)
		copy(dst, cseq)
		dst[len(cseq)] = remaining
		*res = append(*res, dst)
		//return
	}
	for i := start; i * i <= remaining; i++ {
		if remaining % i != 0 {
			continue
		}
		cseq = append(cseq, i)
		factorialCombinationHelper(res, cseq, n, remaining / i, i)
		cseq = cseq[:len(cseq) - 1]
	}
}

