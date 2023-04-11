package medium

func isRobotBounded(instructions string) bool {
	direc := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direcIndex := 0
	x, y := 0, 0
	n := len(instructions)
	for i := 0; i < n; i++ {
		instruction := instructions[i]
		if instruction == 'G' {
			x += direc[direcIndex][0]
			y += direc[direcIndex][1]
		} else if instruction == 'L' {
			direcIndex += 3
			direcIndex %= 4
		} else {
			direcIndex++
			direcIndex %= 4
		}
	}
	return direcIndex != 0 || (x == 0 && y == 0)
}
