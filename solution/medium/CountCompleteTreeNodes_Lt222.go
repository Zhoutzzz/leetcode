package medium

import (
	"leetcode/solution/common"
)

func countNodes(root *common.TreeNode) int {
	return compute(root, 0)
}

func compute(node *common.TreeNode, rem int) int {
	if node == nil {
		return rem
	}
	rem = compute(node.Left, rem)
	rem = compute(node.Right, rem)
	rem++
	return rem
}
