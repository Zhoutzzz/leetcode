package easy

import "leetcode/solution/common"

func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	prev, cur := 0, 0
	for i := 2; i < costLen; i++ {
		prev, cur = cur, common.Min(cur+cost[i-1], prev+cost[i-2])
	}
	return cur
}
