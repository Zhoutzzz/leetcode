package easy

import "sort"

func minSubsequence(nums []int) (res []int) {
	max := 0
	for _, num := range nums {
		max += num
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	greatest := 0
	for i := len(nums) - 1; i >= 0 && max-greatest >= greatest; i-- {
		greatest += nums[i]
		res = append(res, nums[i])
	}
	return
}
