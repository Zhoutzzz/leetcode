package medium

import (
	"leetcode/solution/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxLeaves = 0

func DeepestLeavesSum(root *common.TreeNode) (ans int) {
	var dfs func(node *common.TreeNode, curLeaves int)
	dfs = func(node *common.TreeNode, curLeaves int) {
		if node == nil {
			return
		}
		if curLeaves > maxLeaves {
			ans = node.Val
			maxLeaves = curLeaves
		} else if curLeaves == maxLeaves {
			ans += node.Val
		}
		dfs(node.Left, curLeaves+1)
		dfs(node.Right, curLeaves+1)
	}
	dfs(root, maxLeaves)
	return
}
