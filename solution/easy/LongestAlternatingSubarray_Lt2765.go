package easy

func alternatingSubarray(nums []int) int {
	res := -1
	n := len(nums)
	for firstIndex := 0; firstIndex < n; firstIndex++ {
		for i := firstIndex + 1; i < n; i++ {
			length := i - firstIndex + 1
			if nums[i]-nums[firstIndex] == (length-1)%2 {
				res = max(res, length)
			} else {
				break
			}
		}
	}
	return res
}
