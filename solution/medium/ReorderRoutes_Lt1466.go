package medium

func dfs1466(x, parent int, e [][][]int) int {
	res := 0
	for _, edge := range e[x] {
		if edge[0] == parent {
			continue
		}
		res += edge[1] + dfs1466(edge[0], x, e)
	}
	return res
}

func minReorder(n int, connections [][]int) int {
	e := make([][][]int, n)
	for _, edge := range connections {
		e[edge[0]] = append(e[edge[0]], []int{edge[1], 1})
		e[edge[1]] = append(e[edge[1]], []int{edge[0], 0})
	}
	return dfs1466(0, -1, e)
}
