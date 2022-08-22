package medium

import (
	"leetcode/solution/common"
	"strconv"
)

func calDepth(root *common.TreeNode) (h int) {
	if root.Left != nil {
		h = calDepth(root.Left) + 1
	} else if root.Right != nil {
		h = common.Max(h, calDepth(root.Right)+1)
	}
	return h
}

func printTree(root *common.TreeNode) (ans [][]string) {
	height := calDepth(root)
	var dfs func(node *common.TreeNode, r, c int)
	dfs = func(node *common.TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<height-r-1)
		} else if node.Right != nil {
			dfs(node.Left, r+1, c+1<<height-r-1)
		}
	}
	ans = make([][]string, height+1)
	for i := range ans {
		ans[i] = make([]string, (1<<height+1)-1)
	}

	dfs(root, 0, (len(ans[0])-1)>>1)

	return
}
