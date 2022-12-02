package medium

func minOperations(boxes string) []int {
	left := int(boxes[0] - '0')
	right := 0
	operations := 0
	n := len(boxes)
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			right++
			operations += i
		}
	}
	ans := make([]int, n)
	ans[0] = operations
	for i := 1; i < n; i++ {
		operations += left - right
		if boxes[i] == '1' {
			left++
			right--
		}
		ans[i] = operations
	}
	return ans
}
