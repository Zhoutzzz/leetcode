package medium

import "leetcode/solution/common"

func getSmallestString(n, k int) string {
	s := []byte{}
	for i := 1; i <= n; i++ {
		lower := common.Max(1, k-(n-i)*26)
		k -= lower
		s = append(s, 'a'+byte(lower)-1)
	}
	return string(s)
}
