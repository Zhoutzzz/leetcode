package medium

import "sort"

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	total := int64(0) // 豆子总数
	for _, bean := range beans {
		total += int64(bean)
	}
	res := total // 最少需要移除的豆子数
	for i := 0; i < n; i++ {
		res = min(res, total-int64(beans[i])*int64(n-i))
	}
	return res
}
