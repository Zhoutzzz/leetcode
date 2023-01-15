package easy

import "leetcode/solution/common"

func minMaxGame(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	newNums := make([]int, n/2)
	for i := 0; i < n/2; i++ {
		if i%2 == 0 {
			newNums[i] = common.Min(nums[i*2], nums[i*2+1])
		} else {
			newNums[i] = common.Max(nums[i*2], nums[i*2+1])
		}
	}
	return minMaxGame(newNums)
}
