package medium

import "strconv"

func maximumSwap(num int) int {
	ans := num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			v, _ := strconv.Atoi(string(s))
			ans = max(ans, v)
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}
