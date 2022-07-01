package main

import "sort"

func findContentChildren(g []int, s []int) (res int) {
	sort.Ints(g)
	sort.Ints(s)
	sLen, gLen := len(s), len(g)
	for i, j := 0, 0; i < gLen && j < sLen; i++ {
		for j < sLen && g[i] > s[j] {
			j++
		}
		if j < sLen {
			res++
			j++
		}
	}
	return
}
