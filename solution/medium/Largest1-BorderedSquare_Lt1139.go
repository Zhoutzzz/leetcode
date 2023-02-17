package medium

import "leetcode/solution/common"

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	left := make([][]int, m+1)
	up := make([][]int, m+1)
	for i := range left {
		left[i] = make([]int, n+1)
		up[i] = make([]int, n+1)
	}
	maxBorder := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				left[i][j] = left[i][j-1] + 1
				up[i][j] = up[i-1][j] + 1
				border := common.Min(left[i][j], up[i][j])
				for left[i-border+1][j] < border || up[i][j-border+1] < border {
					border--
				}
				maxBorder = common.Max(maxBorder, border)
			}
		}
	}
	return maxBorder * maxBorder
}
