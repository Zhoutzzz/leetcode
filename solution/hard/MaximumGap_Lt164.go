package hard

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	prev := nums[0]
	maxGap := 0

	for i := 1; i < len(nums); i++ {
		cur := nums[i] - prev
		if cur < 0 {
			cur *= -1
		}
		maxGap = max(cur, maxGap)
		prev = nums[i]
	}

	return maxGap
}
