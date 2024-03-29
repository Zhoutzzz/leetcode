package medium

import (
	"leetcode/solution/common"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points); i++ {
		ans = common.Max(ans, points[i][0]-points[i-1][0])
	}
	return
}
