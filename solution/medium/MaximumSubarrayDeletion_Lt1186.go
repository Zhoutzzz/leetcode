package medium

import "leetcode/solution/common"

func maximumSum(arr []int) int {
	dp0, dp1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		dp0, dp1 = common.Max(dp0, 0)+arr[i], common.Max(dp1+arr[i], dp0)
		res = common.Max(res, common.Max(dp0, dp1))
	}
	return res
}
