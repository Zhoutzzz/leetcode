package medium

import "leetcode/solution/common"

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, g := range grid {
		for j, x := range g {
			if i > 0 {
				f[i][j] = common.Max(f[i][j], f[i-1][j])
			}
			if j > 0 {
				f[i][j] = common.Max((f[i][j], f[i][j-1])
			}
			f[i][j] += x
		}
	}
	return f[m-1][n-1]
}
