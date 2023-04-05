package easy

import "leetcode/solution/common"

func commonFactors(a int, b int) int {
	g := common.Gcd(a, b)
	ans := 0
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g {
				ans++
			}
		}
	}
	return ans
}
