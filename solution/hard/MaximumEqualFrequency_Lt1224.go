package hard

import "leetcode/solution/common"

func maxEqualFreq(nums []int) (ans int) {
	count := make(map[int]int)
	freq := make(map[int]int)
	maxFreq := 0

	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		freq[count[num]]++
		maxFreq = common.Max(maxFreq, count[num])
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = common.Max(ans, i+1)
		}
	}
	return
}
