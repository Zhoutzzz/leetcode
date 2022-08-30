package medium

import "leetcode/solution/common"

func insertIntoMaxTree(root *common.TreeNode, val int) *common.TreeNode {
	if root == nil {
		return &common.TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Right = insertIntoMaxTree(root.Right, val)
	} else {
		return &common.TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	return root
}
