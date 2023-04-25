package easy

import "sort"

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return heights[indices[j]] < heights[indices[i]]
	})
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = names[indices[i]]
	}
	return res
}
