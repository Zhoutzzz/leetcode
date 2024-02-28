package medium

import "leetcode/solution/common"

func minIncrements(n int, cost []int) int {
	ans := 0
	for i := n - 2; i > 0; i -= 2 {
		ans += common.Abs(cost[i] - cost[i+1])
		// 叶节点 i 和 i+1 的双亲节点下标为 i/2（整数除法）
		cost[i/2] = cost[i/2] + max(cost[i], cost[i+1])
	}
	return ans
}
