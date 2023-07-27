package easy

import (
	"leetcode/solution/common"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	res := 0
	for j := 0; j < n; j++ {
		mx := 0
		for i := 0; i < m; i++ {
			mx = common.Max(mx, grid[i][j])
		}
		res += mx
	}
	return res
}
