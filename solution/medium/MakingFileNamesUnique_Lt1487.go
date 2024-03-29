package medium

import "strconv"

func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	index := map[string]int{}
	for p, name := range names {
		i := index[name]
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t
		index[name] = i + 1
		index[t] = 1
	}
	return ans
}
