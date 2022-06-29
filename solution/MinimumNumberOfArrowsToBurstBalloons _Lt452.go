package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(x, y int) bool {
		if points[x][1] < points[y][1] {
			return true
		} else {
			return false
		}
	})

	res := 1
	maxR := points[0][1]

	for _, v := range points {
		if v[0] > maxR {
			maxR = v[1]
			res++
		}
	}
	return res
}
