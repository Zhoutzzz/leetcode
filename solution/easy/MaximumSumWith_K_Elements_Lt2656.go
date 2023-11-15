package easy

func maximizeSum(nums []int, k int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return (2*m + k - 1) * k / 2
}
