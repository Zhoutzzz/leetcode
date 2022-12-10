package hard

import (
	"leetcode/solution/common"
	"sort"
)

func maxHeight(cuboids [][]int) (ans int) {
	for _, c := range cuboids {
		sort.Ints(c)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] < b[0] || a[0] == b[0] && (a[1] < b[1] || a[1] == b[1] && a[2] < b[2])
	})
	f := make([]int, len(cuboids))
	for i, c2 := range cuboids {
		for j, c1 := range cuboids[:i] {
			if c1[1] <= c2[1] && c1[2] <= c2[2] { // 排序后，c1[0] <= c2[0] 恒成立
				f[i] = common.Max(f[i], f[j]) // c1 可以堆在 c2 上
			}
		}
		f[i] += c2[2]
		ans = common.Max(ans, f[i])
	}
	return
}
