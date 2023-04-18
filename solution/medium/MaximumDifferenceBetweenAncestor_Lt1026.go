package medium

import "leetcode/solution/common"

func dfs(root *TreeNode, mi, ma int) int {
	if root == nil {
		return 0
	}
	diff := common.Max(common.Abs(root.Val-mi), common.Abs(root.Val-ma))
	mi, ma = common.Min(mi, root.Val), common.Max(ma, root.Val)
	diff = common.Max(diff, dfs(root.Left, mi, ma))
	diff = common.Max(diff, dfs(root.Right, mi, ma))
	return diff
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}
