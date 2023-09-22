package easy

import "leetcode/solution/common"

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children
	cnt := common.Min(money/7, children)
	money -= cnt * 7
	children -= cnt
	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		cnt -= 1
	}
	return cnt
}
