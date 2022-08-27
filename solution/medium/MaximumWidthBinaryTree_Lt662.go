package medium

import "leetcode/solution/common"

type pair struct {
	index int
	node  *common.TreeNode
}

func WidthOfBinaryTree(root *common.TreeNode) (ans int) {
	q := []pair{{
		index: 1,
		node:  root,
	}}
	tmpq := []pair{}
	for len(q) != 0 {
		p := q[0]
		ans = common.Max(ans, q[len(q)-1].index-p.index+1)
		q = q[1:]
		if p.node.Left != nil {
			tmpq = append(tmpq, pair{
				p.index << 1,
				p.node.Left,
			})
		}
		if p.node.Right != nil {
			tmpq = append(tmpq, pair{
				p.index*2 + 1,
				p.node.Right,
			})
		}
		if len(q) == 0 {
			q = tmpq
			tmpq = []pair{}

		}
	}
	return
}
