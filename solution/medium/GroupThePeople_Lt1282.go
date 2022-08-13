package medium

func groupThePeople(groupSizes []int) (ans [][]int) {
	groupMap := make(map[int][][]int)
	for i, v := range groupSizes {
		if _, ok := groupMap[v]; !ok {
			groupMap[v] = [][]int{{}}
		}
		for gi, groups := range groupMap[v] {
			if len(groups) < v {
				groupMap[v][gi] = append(groups, i)
			} else if gi == len(groupMap[v])-1 {
				groupMap[v] = append(groupMap[v], []int{i})
			}
		}
	}

	for _, v := range groupMap {
		for _, v1 := range v {
			ans = append(ans, v1)
		}
	}

	return
}
