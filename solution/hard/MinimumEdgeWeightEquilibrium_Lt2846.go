package hard

const (
	W = 26
)

func find(uf []int, i int) int {
	if uf[i] == i {
		return i
	}
	uf[i] = find(uf, uf[i])
	return uf[i]
}

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	m := len(queries)
	neighbors := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		neighbors[i] = map[int]int{}
	}
	for _, edge := range edges {
		neighbors[edge[0]][edge[1]] = edge[2]
		neighbors[edge[1]][edge[0]] = edge[2]
	}
	queryArr := make([][][2]int, n)
	for i := 0; i < m; i++ {
		queryArr[queries[i][0]] = append(queryArr[queries[i][0]], [2]int{queries[i][1], i})
		queryArr[queries[i][1]] = append(queryArr[queries[i][1]], [2]int{queries[i][0], i})
	}

	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, W+1)
	}
	visited, uf, lca := make([]int, n), make([]int, n), make([]int, n)
	var tarjan func(int, int)
	tarjan = func(node, parent int) {
		if parent != -1 {
			copy(count[node], count[parent])
			count[node][neighbors[node][parent]]++
		}
		uf[node] = node
		for child, _ := range neighbors[node] {
			if child == parent {
				continue
			}
			tarjan(child, node)
			uf[child] = node
		}
		for _, query := range queryArr[node] {
			node1, index := query[0], query[1]
			if node != node1 && visited[node1] == 0 {
				continue
			}
			lca[index] = find(uf, node1)
		}
		visited[node] = 1
	}
	tarjan(0, -1)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		totalCount, maxCount := 0, 0
		for j := 1; j <= W; j++ {
			t := count[queries[i][0]][j] + count[queries[i][1]][j] - 2*count[lca[i]][j]
			maxCount, totalCount = max(maxCount, t), totalCount+t
		}
		res[i] = totalCount - maxCount
	}
	return res
}
