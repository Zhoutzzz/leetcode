package medium

import "leetcode/solution/common"

const Inf = 0x3f3f3f3f

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return grid[i][j]
		}
		if memo[i][j] >= 0 {
			return memo[i][j]
		}
		res := Inf
		for k := 0; k < n; k++ {
			res = common.Min(res, dfs(i-1, k)+moveCost[grid[i-1][k]][j]+grid[i][j])
		}
		memo[i][j] = res
		return res
	}
	res := Inf
	for j := 0; j < n; j++ {
		res = common.Min(res, dfs(m-1, j))
	}
	return res
}
