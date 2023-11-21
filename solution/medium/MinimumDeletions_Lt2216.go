package medium

func minDeletion(nums []int) int {
	n, ans, check := len(nums), 0, true
	for i := 0; i+1 < n; i++ {
		if nums[i] == nums[i+1] && check {
			ans++
		} else {
			check = !check
		}
	}
	if (n-ans)%2 != 0 {
		ans++
	}
	return ans
}
