package dp

/**
 * 63. Unique Paths II

 You are given an m x n integer array grid. There is a robot initially located at 
 the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right 
 corner (i.e., grid[m - 1][n - 1]). 
 The robot can only move either down or right at any point in time.

An obstacle and space are marked as 1 or 0 respectively in grid. A path that the 
robot takes cannot include any square that is an obstacle.

Return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer will be less than or equal to 2 * 10^9.

Ex 1:
Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Explanation: There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right

Ex 2:
Input: obstacleGrid = [[0,1],[0,0]]
Output: 1

Constraints:

m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] is 0 or 1.
 */

// Here Time:O(M*N), space: O(M*N)
func uniquePathsWithObstaclesDP(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	// 1. Allocate the 2D DP table matching your Java allocation
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 2. Initialize the starting position
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}

	// 3. Initialize the first row
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
		
		if dp[0][i-1] == 0 {
			dp[0][i] = 0
		}
	}

	// 4. Initialize the first column
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
		
		if dp[i-1][0] == 0 {
			dp[i][0] = 0
		}
	}

	// 5. Fill the rest of the DP table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// Here Time:O(M*N), space: O(N) -- improved from previous implementation in terms of space complexitiy.
func uniquePathsWithObstaclesDPE(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	// 1D DP slice representing the current row's path counts
	dp := make([]int, n)

	// Base case: check if the starting position is blocked
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				// If there's an obstacle, 0 paths can pass through this cell
				dp[j] = 0
			} else if j > 0 {
				// dp[j] is the value from the row above
				// dp[j-1] is the value from the left cell in the current row
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}


