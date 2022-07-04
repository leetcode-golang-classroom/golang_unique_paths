package sol

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
	}
	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}
	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[m-1][n-1]
}
