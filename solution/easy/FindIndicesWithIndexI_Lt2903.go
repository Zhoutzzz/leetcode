func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if j - i >= indexDifference && abs(nums[j] - nums[i]) >= valueDifference {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}