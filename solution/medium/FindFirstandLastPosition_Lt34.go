package medium

import "leetcode/solution/common"

func searchRange(nums []int, target int) (res []int) {
	if target > nums[len(nums)-1] || target < nums[0] {
		return append(res, -1, -1)
	}

	res = append(res, common.BinarySearch(nums, 0, len(nums)-1, target))
	res = append(res, common.BinarySearch(nums, res[0], len(nums)-1, target))

	return res
}
