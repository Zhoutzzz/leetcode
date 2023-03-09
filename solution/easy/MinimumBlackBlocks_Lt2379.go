package easy

import "leetcode/solution/common"

func minimumRecolors(blocks string, k int) int {
	res := k
	left, whites := 0, 0
	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			whites++
		}
		if right < k-1 {
			continue
		}
		res = common.Min(res, whites)
		if blocks[left] == 'W' {
			whites--
		}
		left++
	}
	return res
}
