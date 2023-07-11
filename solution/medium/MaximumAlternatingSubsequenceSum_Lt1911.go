package medium

import "leetcode/solution/common"

func maxAlternatingSum(nums []int) int64 {
	even, odd := nums[0], 0
	for i := 1; i < len(nums); i++ {
		even = common.Max(even, odd+nums[i])
		odd = common.Max(odd, even-nums[i])
	}
	return int64(even)
}
