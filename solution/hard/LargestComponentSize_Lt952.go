package hard

import "leetcode/solution/common"

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unionFind{parent, make([]int, n)}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func largestComponentSize(nums []int) (ans int) {
	m := 0
	for _, num := range nums {
		m = common.Max(m, num)
	}
	uf := newUnionFind(m + 1)
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				uf.merge(num, i)
				uf.merge(num, num/i)
			}
		}
	}
	cnt := make([]int, m+1)
	for _, num := range nums {
		rt := uf.find(num)
		cnt[rt]++
		ans = common.Max(ans, cnt[rt])
	}
	return
}
