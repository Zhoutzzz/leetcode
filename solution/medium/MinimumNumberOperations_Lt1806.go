package medium

import "leetcode/solution/common"

func reinitializePermutation(n int) (step int) {
	target := make([]int, n)
	for i := range target {
		target[i] = i
	}
	perm := append([]int(nil), target...)
	for {
		step++
		arr := make([]int, n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if common.Equal(perm, target) {
			return
		}
	}
}
