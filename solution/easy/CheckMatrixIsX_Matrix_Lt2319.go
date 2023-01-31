package easy

func checkXMatrix(grid [][]int) bool {
	for i, row := range grid {
		for j, x := range row {
			if i == j || i+j == len(grid)-1 {
				if x == 0 {
					return false
				}
			} else if x != 0 {
				return false
			}
		}
	}
	return true
}
