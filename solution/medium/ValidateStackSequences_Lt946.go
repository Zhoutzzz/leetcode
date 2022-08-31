package medium

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	for i, j := 0, 0; i < len(pushed) && j < n; {
		if pushed[i] == popped[j] {
			j++
			pushed = append(pushed[:i], pushed[i+1:]...)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return len(pushed) == 0
}
