package easy

func moveZeroes(nums []int) {
	l := len(nums)
	num := 1
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			nums[num-1] = nums[i]
			num++
		}
	}

	for j := l - 1; j >= num; j-- {
		nums[j-1] = 0
	}
}
