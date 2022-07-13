package medium

import "leetcode/solution/common"

func eraseOverlapIntervals(intervals [][]int) (res int) {
	common.QuickSort(intervals, 0, len(intervals)-1)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			res++
		}
	}

	return
}
