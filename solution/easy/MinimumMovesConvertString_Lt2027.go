package easy

func minimumMoves(s string) (res int) {
	covered := -1
	for i, c := range s {
		if c == 'X' && i > covered {
			res++
			covered = i + 2
		}
	}
	return
}
