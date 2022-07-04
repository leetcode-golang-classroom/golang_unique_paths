# golang_unique_paths

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][0]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return *the number of possible unique paths that the robot can take to reach the bottom-right corner*.

The test cases are generated so that the answer will be less than or equal to `2 * $10^9$`.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

```
Input: m = 3, n = 7
Output: 28

```

**Example 2:**

```
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

```

**Constraints:**

- `1 <= m, n <= 100`

## 解析

給定一個整數 m 還有一個整數 n

假設從座標 0, 0 走到 座標 m-1, n-1 的只能往下走或是往右走

求從 座標 0, 0 走到 座標 m-1, n-1 的所有走法的方法數

思考如果要從 座標 0, 0 走到 座標 m-1, n-1  

不管如何走, 最後都是 要走 m-1 個單位往下走, n-1  個單位往右走

所以相當於是在找 m-1 個往走與 n-1 個往右走的排列並且去掉重複的組合

另一個思考方向是考慮要走到 m-1, n-1

由於只能往下跟往右

定義 dp[i,j] = 從 0,0 走到 i-1, j-1 的方法數

則有以下關係式 

dp[i, j] = dp[i-1, j] + dp[i, j-1]

走到 i,j 的方法數 = 走到 i-1, j 的方法數(最後一步往下) + 走到 i, j-1 的方法數(最後一步往右)  

![](https://i.imgur.com/4NPbWZC.png)

時間複雜度 O(m*n)

空間複雜度 O(m*n)

## 程式碼
```go
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

```

## 困難點

1. 要看出每個位置走法數目的遞迴關係

## Solve Point

- [x]  建立一個 m by n 矩陣 用來紀錄已經算過的結果, 初始化 row = 0 或 col = 0 的 cell 值 = 1
- [x]  透過 dp[i][j] = dp[i-1][j] + dp[i][j-1] 來從推算每個值
- [x]  回傳 dp[m-1][n-1]