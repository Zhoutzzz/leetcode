package easy

import "sort"

func arrayRankTransform(arr []int) []int {
	maps := make(map[int]int)
	dst := make([]int, len(arr))
	copy(dst, arr)
	sort.Ints(dst)
	for _, v := range dst {
		if _, ok := maps[v]; !ok {
			maps[v] = len(maps) + 1
		}
	}

	for i, v := range arr {
		arr[i] = maps[v]
	}

	return arr
}
