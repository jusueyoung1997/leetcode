package dp

import (
	"math"
)

// 264. 丑数 II
func nthUglyNumber(n int) int {
	dp := []int{1}
	p1, p2, p3, cnt := 0, 0, 0, 0

	for cnt < n {
		v1, v2, v3, ans := dp[p1]*2, dp[p2]*3, dp[p3]*5, -1
		if v1 <= v2 && v1 <= v3 {
			p1++
			ans = v1
		} else if v2 <= v1 && v2 <= v3 {
			p2++
			ans = v2
		} else {
			p3++
			ans = v3
		}
		if dp[len(dp)-1] != ans {
			dp = append(dp, ans)
			cnt++
		}
	}
	return dp[n-1]
}

// 983. 最低票价
func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n+1)
	duration := []int{1, 7, 30}

	for i, day := range days {
		dp[i+1] = math.MaxInt
		for j, cost := range costs {
			k := i
			for k >= 0 && days[k]+duration[j] > day {
				k--
			}
			dp[i+1] = min(dp[i+1], cost+dp[k+1])
		}
	}
	return dp[n]
}

// 32. 最长有效括号
func longestValidParentheses(s string) int {
	n := len(s)
	dp, ans := make([]int, n+1), 0

	for i, c := range s {
		if c == '(' {
			dp[i+1] = 0
			continue
		}

		l := i - dp[i] - 1
		if l < 0 || s[l] == ')' {
			dp[i+1] = 0
			continue
		}

		dp[i+1] = dp[i] + 2
		if l-1 >= 0 {
			dp[i+1] += dp[l]
		}

		ans = max(ans, dp[i+1])
	}
	return ans
}

// 10. 正则表达式匹配
func isMatch(s string, p string) bool {
	n, m := len(p), len(s)
	dp := make([][]bool, n+1)
	for i := range n + 1 {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true

	for i := range len(p) {
		if p[i] == '*' {
			dp[i+1][0] = dp[i-1][0]
		}

		for j := range len(s) {
			if p[i] == '*' {
				dp[i+1][j+1] = dp[i-1][j+1] || ((s[j] == p[i-1] || p[i-1] == '.') && (dp[i+1][j] || dp[i][j]))
			} else if p[i] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = (s[j] == p[i]) && dp[i][j]
			}
		}
	}
	return dp[n][m]
}

// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		leftup := dp[0]
		for j := 0; j < m; j++ {
			backup := dp[j+1]
			if text2[j] == text1[i] {
				dp[j+1] = max(leftup+1, dp[j+1])
			}
			dp[j+1] = max(dp[j+1], dp[j])
			leftup = backup
		}
	}
	return dp[m]
}

// 516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		leftup := dp[0]
		for j := 0; j < n; j++ {
			backup := dp[j+1]
			if s[i] == s[n-1-j] {
				dp[j+1] = max(leftup+1, dp[j+1])
			}
			dp[j+1] = max(dp[j+1], dp[j])
			leftup = backup
		}
	}
	return dp[n]
}

// LCR 112. 矩阵中的最长递增路径
func longestIncreasingPath(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	axis := []int{1, 0, -1, 0, 1}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 从(x, y)出发的最长路径
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if dp[x][y] != 0 {
			return dp[x][y]
		}
		dp[x][y] = 1
		for i := 0; i < 4; i++ {
			dx, dy := x+axis[i], y+axis[i+1]
			if dx < 0 || dx >= n || dy < 0 || dy >= m || matrix[dx][dy] <= matrix[x][y] {
				continue
			}
			dp[x][y] = max(dp[x][y], dfs(dx, dy)+1)
		}
		return dp[x][y]
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
