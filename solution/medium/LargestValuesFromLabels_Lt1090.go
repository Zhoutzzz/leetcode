package medium

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})

	ans, choose := 0, 0
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		label := labels[idx[i]]
		if cnt[label] == useLimit {
			continue
		}
		choose++
		ans += values[idx[i]]
		cnt[label]++
		if choose == numWanted {
			break
		}
	}
	return ans
}
