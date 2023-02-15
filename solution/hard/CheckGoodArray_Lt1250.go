package hard

import "leetcode/solution/common"

func isGoodArray(nums []int) bool {
	g := 0
	for _, x := range nums {
		g = common.Gcd(g, x)
		if g == 1 {
			return true
		}
	}
	return false
}
