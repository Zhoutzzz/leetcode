package easy

func validPath(n int, edges [][]int, source int, destination int) bool {

	bset := make([]int, n+1)

	var findInSet func(a int) int
	findInSet = func(a int) int {
		if bset[a] == 0 {
			return a
		}
		bset[a] = findInSet(bset[a])
		return bset[a]
	}

	var unionSet func(a, b int)
	unionSet = func(a, b int) {
		a1 := findInSet(a)
		b1 := findInSet(b)
		if a1 != b1 {
			bset[a1] = b1
		}
	}

	for _, e := range edges {
		a, b := e[0]+1, e[1]+1
		unionSet(a, b)
	}

	return findInSet(source+1) == findInSet(destination+1)

}
