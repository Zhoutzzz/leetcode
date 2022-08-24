package easy

func canBeEqual(target []int, arr []int) bool {
	maps := make(map[int]int, len(target))
	for i, v := range target {
		maps[v]++
		maps[arr[i]]--
	}
	for _, v := range maps {
		if v != 0 {
			return false
		}
	}
	return true
}
