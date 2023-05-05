package easy

func hardestWorker(n int, logs [][]int) int {
	ans, maxCost := logs[0][0], logs[0][1]
	for i := 1; i < len(logs); i++ {
		idx := logs[i][0]
		cost := logs[i][1] - logs[i-1][1]
		if cost > maxCost || (cost == maxCost && idx < ans) {
			ans = idx
			maxCost = cost
		}
	}
	return ans
}
