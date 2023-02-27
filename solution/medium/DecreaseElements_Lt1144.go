package medium

import "leetcode/solution/common"

func movesToMakeZigzag(nums []int) int {
	help := func(pos int) int {
		res := 0
		for i := pos; i < len(nums); i += 2 {
			a := 0
			if i-1 >= 0 {
				a = common.Min(a, nums[i]-nums[i-1]+1)
			}
			if i+1 < len(nums) {
				a = common.Max(a, nums[i]-nums[i+1]+1)
			}
			res += a
		}
		return res
	}

	return common.Min(help(0), help(1))
}
