package easy

import "leetcode/solution/common"

func distanceBetweenBusStops(distance []int, start int, destination int) (res int) {
	if start > destination {
		start, destination = destination, start
	}
	a, b := 0, 0
	for i, v := range distance {
		if start <= i && i < destination {
			a += v
		} else {
			b += v
		}
	}
	return common.Min(a, b)
}
