package medium

import "leetcode/solution/common"

func maximumScore(a, b, c int) int {
	sum := a + b + c
	maxVal := common.Max(common.Max(a, b), c)
	if sum < maxVal*2 {
		return sum - maxVal
	} else {
		return sum / 2
	}
}
