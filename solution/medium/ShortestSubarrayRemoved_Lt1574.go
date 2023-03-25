package medium

import "leetcode/solution/common"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	res := j
	for i := 0; i < n; i++ {
		for j < n && arr[j] < arr[i] {
			j++
		}
		res = common.Min(res, j-i-1)
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}
