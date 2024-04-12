package easy

func findChampion(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		sum := 0
		for _, val := range grid[i] {
			sum += val
		}
		if sum == n-1 {
			return i
		}
	}
	return -1
}
