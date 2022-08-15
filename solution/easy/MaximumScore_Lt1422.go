package easy

import (
	"leetcode/solution/common"
	"strings"
)

func maxScore(s string) (ans int) {
	score := int('1'-s[0]) + strings.Count(s[1:], "1")
	ans = score
	for _, c := range s[1 : len(s)-1] {
		if c == '0' {
			score++
		} else {
			score--
		}
		ans = common.Max(ans, score)
	}
	return
}
