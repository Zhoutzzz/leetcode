package easy

import "leetcode/solution/common"

func minStartValue(nums []int) (ans int) {
	min, normal := 0, 0
	for _, v := range nums {
		normal += v
		min = common.Min(min, normal)
	}

	return -min + 1
}
