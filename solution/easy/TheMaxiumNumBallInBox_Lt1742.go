package easy

import "leetcode/solution/common"

func countBalls(lowLimit, highLimit int) (ans int) {
	count := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		count[sum]++
		ans = common.Max(ans, count[sum])
	}
	return
}
