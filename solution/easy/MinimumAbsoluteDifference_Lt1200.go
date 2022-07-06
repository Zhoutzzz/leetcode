package easy

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) (res [][]int) {
	arrLen := len(arr)
	sort.Ints(arr)

	for i, abs := 0, math.MaxInt; i < arrLen-1; i++ {
		if curAbs := arr[i+1] - arr[i]; curAbs < abs {
			abs = curAbs
			res = [][]int{{arr[i], arr[i+1]}}
		} else if curAbs == abs {
			res = append(res, []int{arr[i], arr[i+1]})

		}
	}

	return res
}
