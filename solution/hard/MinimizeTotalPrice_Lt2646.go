package hard

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	next := make([][]int, n)
	for _, edge := range edges {
		next[edge[0]] = append(next[edge[0]], edge[1])
		next[edge[1]] = append(next[edge[1]], edge[0])
	}

	count := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(node, parent, end int) bool {
		if node == end {
			count[node]++
			return true
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			if dfs(child, node, end) {
				count[node]++
				return true
			}
		}
		return false
	}
	for _, trip := range trips {
		dfs(trip[0], -1, trip[1])
	}

	var dp func(int, int) []int
	dp = func(node, parent int) []int {
		res := []int{
			price[node] * count[node], price[node] * count[node] / 2,
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			v := dp(child, node)
			// node 没有减半，因此可以取子树的两种情况的最小值
			// node 减半，只能取子树没有减半的情况
			res[0], res[1] = res[0]+min(v[0], v[1]), res[1]+v[0]
		}
		return res
	}
	res := dp(0, -1)
	return min(res[0], res[1])
}
