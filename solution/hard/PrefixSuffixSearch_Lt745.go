package hard

import (
	"leetcode/solution/common"
)

type WordFilter struct {
	children map[string]WordFilter
	weight   map[string]int
}

const PACKING string = "##"

func Constructor(words []string) WordFilter {
	root := WordFilter{children: make(map[string]WordFilter)}
	for i, v := range words {
		cur := root
		wordLen := len(v)
		for j := range v {
			tmp := cur
			for k := j; k < wordLen; k++ {
				thisChar := string(v[k]) + "#"
				if _, ok := tmp.children[thisChar]; !ok {
					tmp.children[thisChar] = WordFilter{children: make(map[string]WordFilter), weight: make(map[string]int)}
				}
				tmp = tmp.children[thisChar]
				tmp.weight[PACKING] = i
			}
			tmp = cur
			for k := j; k < wordLen; k++ {
				thisChar := "#" + string(v[wordLen-k-1])
				if _, ok := tmp.children[thisChar]; !ok {
					tmp.children[thisChar] = WordFilter{children: make(map[string]WordFilter), weight: make(map[string]int)}
				}
				tmp = tmp.children[thisChar]
				tmp.weight[PACKING] = i
			}
			thisChar := string(v[j]) + string(v[wordLen-j-1])
			if _, ok := cur.children[thisChar]; !ok {
				cur.children[thisChar] = WordFilter{children: make(map[string]WordFilter), weight: make(map[string]int)}
			}
			cur = cur.children[thisChar]
			cur.weight[PACKING] = i
		}
	}

	return root
}

func (this *WordFilter) F(pref string, suff string) int {
	maxLen := common.Max(len(pref), len(suff))
	cur := *this
	for i := 0; i < maxLen; i++ {
		curStr := ""
		if i < len(pref) {
			curStr += string(pref[i])
		} else {
			curStr += "#"
		}
		if i < len(suff) {
			curStr += string(suff[len(suff)-i-1])
		} else {
			curStr += "#"
		}
		if j, ok := cur.children[curStr]; !ok {
			println(j.weight[PACKING])
			return -1
		} else {
			cur = cur.children[curStr]
		}
	}
	return cur.weight[PACKING]
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
