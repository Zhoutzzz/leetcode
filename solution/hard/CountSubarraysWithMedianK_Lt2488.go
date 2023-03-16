package hard

import "leetcode/solution/common"

func countSubarrays(nums []int, k int) int {
	kIndex := -1
	for i, num := range nums {
		if num == k {
			kIndex = i
			break
		}
	}
	ans := 0
	counts := map[int]int{}
	counts[0] = 1
	sum := 0
	for i, num := range nums {
		sum += common.Sign(num - k)
		if i < kIndex {
			counts[sum]++
		} else {
			prev0 := counts[sum]
			prev1 := counts[sum-1]
			ans += prev0 + prev1
		}
	}
	return ans
}
