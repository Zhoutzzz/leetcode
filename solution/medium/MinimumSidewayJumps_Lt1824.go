package medium

import (
	"leetcode/solution/common"
	"math"
)

func minSideJumps(obstacles []int) int {
	d := [3]int{1, 0, 1}
	for _, x := range obstacles[1:] {
		minCnt := math.MaxInt / 2
		for j := 0; j < 3; j++ {
			if j == x-1 {
				d[j] = math.MaxInt / 2
			} else {
				minCnt = common.Min(minCnt, d[j])
			}
		}
		for j := 0; j < 3; j++ {
			if j != x-1 {
				d[j] = common.Min(d[j], minCnt+1)
			}
		}
	}
	return common.Min(common.Min(d[0], d[1]), d[2])
}
