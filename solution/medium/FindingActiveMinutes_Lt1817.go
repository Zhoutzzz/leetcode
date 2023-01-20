package medium

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	mp := map[int]map[int]bool{}
	for _, p := range logs {
		id, t := p[0], p[1]
		if mp[id] == nil {
			mp[id] = map[int]bool{}
		}
		mp[id][t] = true
	}
	ans := make([]int, k+1)
	for _, m := range mp {
		ans[len(m)]++
	}
	return ans[1:]
}
