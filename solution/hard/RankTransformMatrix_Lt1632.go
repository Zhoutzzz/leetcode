package hard

import (
	"leetcode/solution/common"
	"sort"
)

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	type pair struct{ i, j int }
	d := map[int][]pair{}
	for i, row := range matrix {
		for j, v := range row {
			d[v] = append(d[v], pair{i, j})
		}
	}
	rowMax := make([]int, m)
	colMax := make([]int, n)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	vs := []int{}
	for v := range d {
		vs = append(vs, v)
	}
	sort.Ints(vs)
	uf := common.NewUnionFind(m + n)
	rank := make([]int, m+n)
	for _, v := range vs {
		ps := d[v]
		for _, p := range ps {
			uf.Union(p.i, p.j+m)
		}
		for _, p := range ps {
			i, j := p.i, p.j
			rank[uf.Find(i)] = common.Max(rank[uf.Find(i)], common.Max(rowMax[i], colMax[j]))
		}
		for _, p := range ps {
			i, j := p.i, p.j
			ans[i][j] = 1 + rank[uf.Find(i)]
			rowMax[i], colMax[j] = ans[i][j], ans[i][j]
		}
		for _, p := range ps {
			uf.Reset(p.i)
			uf.Reset(p.j + m)
		}
	}
	return ans
}
