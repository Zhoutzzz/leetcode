package main

import "strconv"

func matrixScore(A [][]int) (res int) {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		if A[i][0] == 0 {
			for j := 0; j < n; j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}

	for j := 1; j < n; j++ {
		count := 0
		for i := 0; i < m; i++ {
			if A[i][j] == 1 {
				count++
			}
		}

		curLen := float64(m) / 2.0

		if float64(count) < curLen {
			for k := 0; k < len(A); k++ {
				A[k][j] = 1 - A[k][j]
			}
		}
	}

	for i := 0; i < len(A); i++ {
		cur := ""
		for j := 0; j < len(A[0]); j++ {
			cur += strconv.Itoa(A[i][j])
		}
		p, _ := strconv.ParseInt(cur, 2, 64)
		res += int(p)
	}

	return res
}
