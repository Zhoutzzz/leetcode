package hard

import (
	"leetcode/solution/common"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	n := len(satisfaction)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	sort.Ints(satisfaction)
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = dp[i-1][j-1] + satisfaction[i-1]*j
			if j < i {
				dp[i][j] = common.Max(dp[i][j], dp[i-1][j])
			}
			res = common.Max(res, dp[i][j])
		}
	}
	return res
}
