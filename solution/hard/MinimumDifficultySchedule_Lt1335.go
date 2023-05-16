package hard

import "leetcode/solution/common"

func minDifficulty(a []int, d int) int {
	n := len(a)
	if n < d {
		return -1
	}

	f := make([][]int, d)
	f[0] = make([]int, n)
	f[0][0] = a[0]
	for i := 1; i < n; i++ {
		f[0][i] = common.Max(f[0][i-1], a[i])
	}
	for i := 1; i < d; i++ {
		f[i] = make([]int, n)
		type pair struct{ j, mn int }
		st := []pair{} // (下标 j，从 f[i-1][left[j]] 到 f[i-1][j-1] 的最小值)
		for j := i; j < n; j++ {
			mn := f[i-1][j-1]                               // 只有 a[j] 一项工作
			for len(st) > 0 && a[st[len(st)-1].j] <= a[j] { // 向左一直计算到 left[j]
				mn = common.Min(mn, st[len(st)-1].mn)
				st = st[:len(st)-1]
			}
			f[i][j] = mn + a[j] // 从 a[left[j]+1] 到 a[j] 的最大值是 a[j]
			if len(st) > 0 {    // 如果这一段包含 <=left[j] 的工作，那么这一段的最大值必然不是 a[j]
				f[i][j] = common.Min(f[i][j], f[i][st[len(st)-1].j]) // 答案和 f[i][left[j]] 是一样的
			}
			st = append(st, pair{j, mn}) // 注意这里保存的不是 f[i][j]
		}
	}
	return f[d-1][n-1]
}
