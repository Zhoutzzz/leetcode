package hard

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	st1, st2 := []int{}, []int{}

	for i := 0; i < n; i++ {
		v := nums[i]
		for len(st2) > 0 && nums[st2[len(st2)-1]] < v {
			res[st2[len(st2)-1]] = v
			st2 = st2[:len(st2)-1]
		}
		pos := len(st1) - 1
		for pos >= 0 && nums[st1[pos]] < v {
			pos--
		}
		st2 = append(st2, st1[pos+1:]...)
		st1 = append(st1[:pos+1], i)
	}
	return res
}
