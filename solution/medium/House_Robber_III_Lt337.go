package medium

import "leetcode/solution/common"

func rob(root *TreeNode) int {
	val := dfs_337(root)
	return common.Max(val[0], val[1])
}

func dfs_337(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs_337(node.Left), dfs_337(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := common.Max(l[0], l[1]) + common.Max(r[0], r[1])
	return []int{selected, notSelected}
}
