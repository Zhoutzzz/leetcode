package hard

import (
	"leetcode/solution/common"
)

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = common.Max(leftMax, height[left])
		rightMax = common.Max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

type mystr struct {
	length int
}

func (a mystr) Len() int {
	//TODO implement me
	return a.length
}

func (mystr) Less(i, j int) bool {
	//TODO implement me
	return i > j
}

func (mystr) Swap(i, j int) {
	//TODO implement me
	i = j
}

func a() interface{} {
	return nil
}
