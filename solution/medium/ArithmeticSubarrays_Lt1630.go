package medium

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := range l {
		ans[i] = isArithmetic(nums[l[i] : r[i]+1])
	}
	return ans
}

func isArithmetic(nums []int) bool {
	n := len(nums)
	max, min := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if max == min {
		return true
	}
	d := (max - min) / (n - 1)
	if (max - min) != d*(n-1) {
		return false
	}
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for n--; n > 0; n-- {
		min += d
		if _, ok := m[min]; !ok {
			return false
		}
	}
	return true
}
