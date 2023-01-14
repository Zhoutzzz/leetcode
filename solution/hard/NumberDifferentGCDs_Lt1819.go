package hard

import "leetcode/solution/common"

func countDifferentSubsequenceGCDs(nums []int) (ans int) {
	maxVal := 0
	for _, num := range nums {
		maxVal = common.Max(maxVal, num)
	}
	occured := make([]bool, maxVal+1)
	for _, num := range nums {
		occured[num] = true
	}
	for i := 1; i <= maxVal; i++ {
		subGcd := 0
		for j := i; j <= maxVal; j += i {
			if occured[j] {
				if subGcd == 0 {
					subGcd = j
				} else {
					subGcd = common.Gcd(subGcd, j)
				}
				if subGcd == i {
					ans++
					break
				}
			}
		}
	}
	return
}
