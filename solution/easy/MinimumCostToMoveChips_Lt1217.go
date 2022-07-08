package easy

import (
	"leetcode/solution/common"
)

func minCostToMoveChips(position []int) int {
	i, j := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			i++
		} else {
			j++
		}
	}
	return common.Min(i, j)
}
