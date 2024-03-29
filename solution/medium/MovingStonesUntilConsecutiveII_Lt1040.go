package medium

import (
	"leetcode/solution/common"
	"sort"
)

func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)
	if stones[n-1]-stones[0]+1 == n {
		return []int{0, 0}
	}
	ma := common.Max(stones[n-2]-stones[0]+1, stones[n-1]-stones[1]+1) - (n - 1)
	mi := n
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i]+1 <= n {
			j++
		}
		if j-i+1 == n-1 && stones[j]-stones[i]+1 == n-1 {
			mi = common.Min(mi, 2)
		} else {
			mi = common.Min(mi, n-(j-i+1))
		}
	}
	return []int{mi, ma}
}
