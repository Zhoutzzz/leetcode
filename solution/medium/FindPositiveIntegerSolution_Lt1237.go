package medium

func findSolution(customFunction func(int, int) int, z int) (ans [][]int) {
	for x := 1; x <= 1000; x++ {
		for y := 1; y <= 1000; y++ {
			if customFunction(x, y) == z {
				ans = append(ans, []int{x, y})
			}
		}
	}
	return
}
