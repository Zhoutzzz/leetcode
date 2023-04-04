package hard

import "leetcode/solution/common"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = 1e9
		}
	}
	sum := make([]int, n+1)
	for i, s := 0, 0; i < n; i++ {
		d[i][i] = 0
		s += stones[i]
		sum[i+1] = s
	}

	for len := 2; len <= n; len++ {
		for l := 0; l < n && l+len-1 < n; l++ {
			r := l + len - 1
			for p := l; p < r; p += k - 1 {
				d[l][r] = common.Min(d[l][r], d[l][p]+d[p+1][r])
			}
			if (r-l)%(k-1) == 0 {
				d[l][r] += sum[r+1] - sum[l]
			}
		}
	}
	return d[0][n-1]
}
