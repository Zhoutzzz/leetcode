package easy

func oddCells(row int, column int, indices [][]int) (res int) {
	metrics := make([][]int, row)
	for i := range metrics {
		metrics[i] = make([]int, column)
	}
	for _, v := range indices {
		for _, v1 := range metrics {
			v1[v[1]]++
			if v1[v[1]]%2 == 1 {
				res++
			} else if res > 0 {
				res--
			}
		}
		for i := range metrics[v[0]] {
			metrics[v[0]][i]++
			if metrics[v[0]][i]%2 == 1 {
				res++
			} else if res > 0 {
				res--
			}
		}
	}

	return res
}

//func oddCells(m, n int, indices [][]int) int {
//	rows := make([]int, m)
//	cols := make([]int, n)
//	for _, p := range indices {
//		rows[p[0]]++
//		cols[p[1]]++
//	}
//	oddx := 0
//	for _, row := range rows {
//		oddx += row % 2
//	}
//	oddy := 0
//	for _, col := range cols {
//		oddy += col % 2
//	}
//	return oddx*(n-oddy) + (m-oddx)*oddy
//}
