package medium

import "leetcode/solution/common"

func help_1775(h1 [7]int, h2 [7]int, diff int) (res int) {
	h := [7]int{}
	for i := 1; i < 7; i++ {
		h[6-i] += h1[i]
		h[i-1] += h2[i]
	}
	for i := 5; i > 0 && diff > 0; i-- {
		t := common.Min((diff+i-1)/i, h[i])
		res += t
		diff -= t * i
	}
	return res
}

func minOperations_1775(nums1 []int, nums2 []int) (ans int) {
	n, m := len(nums1), len(nums2)
	if 6*n < m || 6*m < n {
		return -1
	}
	var cnt1, cnt2 [7]int
	diff := 0
	for _, i := range nums1 {
		cnt1[i]++
		diff += i
	}
	for _, i := range nums2 {
		cnt2[i]++
		diff -= i
	}
	if diff == 0 {
		return 0
	}
	if diff > 0 {
		return help_1775(cnt2, cnt1, diff)
	}
	return help_1775(cnt1, cnt2, -diff)
}
