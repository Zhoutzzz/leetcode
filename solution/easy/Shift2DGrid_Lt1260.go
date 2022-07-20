package easy

func shiftGrid(grid [][]int, k int) [][]int {
	for i := 0; i < k; i++ {
		prev := grid[0][0]
		for j, row := range grid {
			for i := 0; i < len(row); i++ {
				tmp := grid[j][i]
				grid[j][i] = prev
				prev = tmp
			}
		}
		grid[0][0] = prev
	}
	return grid
}

//func shiftGrid(grid [][]int, k int) [][]int {
//	m, n := len(grid), len(grid[0])
//	ans := make([][]int, m)
//	for i := range ans {
//		ans[i] = make([]int, n)
//	}
//	for i, row := range grid {
//		for j, v := range row {
//			index1 := (i*n + j + k) % (m * n)
//			ans[index1/n][index1%n] = v
//		}
//	}
//	return ans
//}
