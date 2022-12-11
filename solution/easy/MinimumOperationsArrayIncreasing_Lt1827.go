package easy

import "leetcode/solution/common"

func minOperations(nums []int) (ans int) {
	pre := nums[0] - 1
	for _, x := range nums {
		pre = common.Max(pre+1, x)
		ans += pre - x
	}
	return
}
