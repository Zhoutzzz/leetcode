package medium

import "leetcode/solution/common"

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	i, j := 0, 0
	for i < n && j < m {
		v := common.Min(rowSum[i], colSum[j])
		matrix[i][j] = v
		rowSum[i] -= v
		colSum[j] -= v
		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return matrix
}
