package medium

func arrayNesting(nums []int) (max int) {
	n := len(nums)
	for i := range nums {
		cnt := 0
		for nums[i] < n {
			cnt++
			i, nums[i] = nums[i], n
		}
		if cnt > max {
			max = cnt
		}
	}

	return
}
