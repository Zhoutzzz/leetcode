package medium

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, numCourses)
	indgree := make([]int, numCourses)
	isPre := make([][]bool, numCourses)
	for i, _ := range isPre {
		isPre[i] = make([]bool, numCourses)
		g[i] = []int{}
	}
	for _, p := range prerequisites {
		indgree[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, ne := range g[cur] {
			isPre[cur][ne] = true
			for i := 0; i < numCourses; i++ {
				isPre[i][ne] = isPre[i][ne] || isPre[i][cur]
			}
			indgree[ne]--
			if indgree[ne] == 0 {
				q = append(q, ne)
			}
		}
	}
	res := []bool{}
	for _, query := range queries {
		res = append(res, isPre[query[0]][query[1]])
	}
	return res
}
