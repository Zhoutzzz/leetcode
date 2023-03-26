package easy

func findSubarrays(nums []int) bool {
	seen := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
