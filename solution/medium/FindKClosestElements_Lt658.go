package medium

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	j := sort.SearchInts(arr, x)
	i := j - 1
	for ; k > 0; k-- {
		if i < 0 {
			j++
		} else if j >= l || x-arr[i] <= arr[j]-x {
			i--
		} else {
			j++
		}
	}
	return arr[i+1 : j]
}
