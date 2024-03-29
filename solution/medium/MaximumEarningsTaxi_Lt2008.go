package medium

func maxTaxiEarnings(n int, rides [][]int) int64 {
	type pair struct{ s, t int }
	groups := make([][]pair, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], pair{start, end - start + tip})
	}

	f := make([]int64, n+1)
	for i := 2; i <= n; i++ {
		f[i] = f[i-1]
		for _, p := range groups[i] {
			f[i] = max(f[i], f[p.s]+int64(p.t))
		}
	}
	return f[n]
}
