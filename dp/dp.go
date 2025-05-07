package dp

import "math"

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
