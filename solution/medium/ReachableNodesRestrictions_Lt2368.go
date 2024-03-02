package medium

func reachableNodes(n int, edges [][]int, restricted []int) int {
	isRestricted := make([]int, n)
	for _, x := range restricted {
		isRestricted[x] = 1
	}
	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	cnt := 0
	var dfs func(int, int)
	dfs = func(x, f int) {
		cnt++
		for _, y := range g[x] {
			if y != f && isRestricted[y] != 1 {
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)
	return cnt
}
