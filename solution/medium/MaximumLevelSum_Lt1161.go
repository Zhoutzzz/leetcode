package medium

import "leetcode/solution/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *common.TreeNode) int {
	ans, maxSum := 1, root.Val
	q := []*common.TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		tmp := q
		q = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}
