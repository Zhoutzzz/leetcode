package medium

import "leetcode/solution/common"

func spiralMatrix(m int, n int, head *common.ListNode) (res [][]int) {
	res = make([][]int, m)
	for i := 0; i < m; i++ {
		q := make([]int, n)
		for j := 0; j < n; j++ {
			q[j] = -1
		}
		res[i] = q
	}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	line, column, di := 0, 0, 0
	for head != nil {
		res[line][column] = head.Val
		head = head.Next
		if line+dir[di][0] >= m || line+dir[di][0] < 0 || column+dir[di][1] >= n || dir[di][1]+column < 0 || res[dir[di][0]+line][dir[di][1]+column] != -1 {
			di = (di + 1) % 4
		}
		i, j := dir[di][0], dir[di][1]
		line += i
		column += j
	}

	return res
}
