package easy

import "leetcode/solution/common"

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	l := len(s)
	for b, e := 0, len(s)-1; b <= e; {
		if l > ans && s[b] == s[e] {
			ans = common.Max(ans, e-b-1)
			b++
			l--
			e = len(s) - 1
		} else {
			e--
		}
	}

	return ans
}
