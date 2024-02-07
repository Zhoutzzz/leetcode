package medium

func replaceValueInTree(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	root.Val = 0
	for len(q) > 0 {
		var q2 []*TreeNode
		sum := 0
		for _, fa := range q {
			if fa.Left != nil {
				q2 = append(q2, fa.Left)
				sum += fa.Left.Val
			}
			if fa.Right != nil {
				q2 = append(q2, fa.Right)
				sum += fa.Right.Val
			}
		}
		for _, fa := range q {
			childSum := 0
			if fa.Left != nil {
				childSum += fa.Left.Val
			}
			if fa.Right != nil {
				childSum += fa.Right.Val
			}
			if fa.Left != nil {
				fa.Left.Val = sum - childSum
			}
			if fa.Right != nil {
				fa.Right.Val = sum - childSum
			}
		}
		q = q2
	}
	return root
}
