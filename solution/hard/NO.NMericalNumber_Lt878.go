package hard

import "leetcode/solution/common"

const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	l := common.Min(a, b)
	r := n * common.Min(a, b)
	c := a / common.Gcd(a, b) * b
	for l <= r {
		mid := (l + r) / 2
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}
