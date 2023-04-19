package medium

import "leetcode/solution/common"

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		maxValue := arr[i-1]
		for j := i - 1; j >= common.Max(0, i-k); j-- {
			d[i] = common.Max(d[i], d[j]+maxValue*(i-j))
			if j > 0 && arr[j-1] > maxValue {
				maxValue = arr[j-1]
			}
		}
	}
	return d[n]
}
