package hard

import "leetcode/solution/common"

const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	l := min(a, b)
	r := n * min(a, b)
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
