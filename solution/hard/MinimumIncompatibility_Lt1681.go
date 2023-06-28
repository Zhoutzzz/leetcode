package hard

import (
	"leetcode/solution/common"
	"math"
	"math/bits"
)

func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	group := n / k
	inf := math.MaxInt32
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	values := make(map[int]int)

	for mask := 1; mask < (1 << n); mask++ {
		if bits.OnesCount(uint(mask)) != group {
			continue
		}
		minVal := 20
		maxVal := 0
		cur := make(map[int]bool)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if cur[nums[i]] {
					break
				}
				cur[nums[i]] = true
				minVal = common.Min(minVal, nums[i])
				maxVal = common.Max(maxVal, nums[i])
			}
		}
		if len(cur) == group {
			values[mask] = maxVal - minVal
		}
	}

	for mask := 0; mask < (1 << n); mask++ {
		if dp[mask] == inf {
			continue
		}
		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) == 0 {
				seen[nums[i]] = i
			}
		}
		if len(seen) < group {
			continue
		}
		sub := 0
		for _, v := range seen {
			sub |= (1 << v)
		}
		nxt := sub
		for nxt > 0 {
			if val, ok := values[nxt]; ok {
				dp[mask|nxt] = common.Min(dp[mask|nxt], dp[mask]+val)
			}
			nxt = (nxt - 1) & sub
		}
	}
	if dp[(1<<n)-1] < inf {
		return dp[(1<<n)-1]
	}
	return -1
}
