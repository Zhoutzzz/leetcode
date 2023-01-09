package medium

func canChoose(groups [][]int, nums []int) bool {
next:
	for _, g := range groups {
		for len(nums) >= len(g) {
			if equal(nums[:len(g)], g) {
				nums = nums[len(g):]
				continue next
			}
			nums = nums[1:]
		}
		return false
	}
	return true
}
