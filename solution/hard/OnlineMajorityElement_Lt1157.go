package hard

import (
	"math/rand"
	"sort"
	"time"
)

type MajorityChecker struct {
	arr []int
	loc map[int][]int
}

func Constructor1157(arr []int) MajorityChecker {
	rand.Seed(time.Now().UnixNano())
	loc := map[int][]int{}
	for i, x := range arr {
		loc[x] = append(loc[x], i)
	}
	return MajorityChecker{arr, loc}
}

func (mc *MajorityChecker) Query(left, right, threshold int) int {
	length := right - left + 1
	for i := 0; i < 20; i++ {
		x := mc.arr[rand.Intn(right-left+1)+left]
		pos := mc.loc[x]
		occ := sort.SearchInts(pos, right+1) - sort.SearchInts(pos, left)
		if occ >= threshold {
			return x
		}
		if occ*2 >= length {
			break
		}
	}
	return -1
}
