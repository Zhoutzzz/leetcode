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
type CBTInserter struct {
	root      *common.TreeNode
	nodeQueue []*common.TreeNode
}

func Constructor(root *common.TreeNode) CBTInserter {
	nodes := []*common.TreeNode{root}
	return CBTInserter{
		root,
		nodes,
	}
}

func (this *CBTInserter) Insert(val int) int {
	for {
		news := &common.TreeNode{
			Val: val,
		}
		cur := this.nodeQueue[0]
		if cur.Left == nil {
			cur.Left = news
			return cur.Val
		} else if cur.Right == nil {
			cur.Right = news
			this.nodeQueue = append(this.nodeQueue, cur.Left)
			this.nodeQueue = append(this.nodeQueue, cur.Right)
			this.nodeQueue = this.nodeQueue[1:]
			return cur.Val
		} else {
			this.nodeQueue = append(this.nodeQueue, cur.Left)
			this.nodeQueue = append(this.nodeQueue, cur.Right)
			this.nodeQueue = this.nodeQueue[1:]
		}
	}
}

func (this *CBTInserter) Get_root() *common.TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
