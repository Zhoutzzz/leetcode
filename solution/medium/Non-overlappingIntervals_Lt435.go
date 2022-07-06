package medium

func eraseOverlapIntervals(intervals [][]int) (res int) {
	quickSort(intervals, 0, len(intervals)-1)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			res++
		}
	}

	return
}
