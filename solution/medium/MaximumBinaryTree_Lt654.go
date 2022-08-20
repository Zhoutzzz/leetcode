package medium

import (
	"leetcode/solution/common"
	"math"
)

func constructMaximumBinaryTree(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx, maxVal := 0, math.MinInt
	for i, num := range nums {
		if maxVal < common.Max(num, maxVal) {
			maxVal = common.Max(num, maxVal)
			maxIdx = i
		}
	}
	root := common.TreeNode{Val: maxVal}
	left := nums[:maxIdx]
	root.Left = constructMaximumBinaryTree(left)
	right := nums[maxIdx+1:]
	root.Right = constructMaximumBinaryTree(right)
	return &root
}
