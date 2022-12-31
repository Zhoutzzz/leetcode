package easy

import (
	"leetcode/solution/common"
	"sort"
)

func minMovesToSeat(seats, students []int) (ans int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, x := range seats {
		ans += common.Abs(x - students[i])
	}
	return
}
