package medium

import "leetcode/solution/common"

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	dp, mval := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		mval[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		mval[j][j] = arr[j]
		for i := j - 1; i >= 0; i-- {
			mval[i][j] = common.Max(arr[i], mval[i+1][j])
			dp[i][j] = 0x3f3f3f3f
			for k := i; k < j; k++ {
				dp[i][j] = common.Min(dp[i][j], dp[i][k]+dp[k+1][j]+mval[i][k]*mval[k+1][j])
			}
		}
	}
	return dp[0][n-1]
}
