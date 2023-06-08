package hard

import "leetcode/solution/common"

func tilingRectangle(n int, m int) int {
	ans := common.Max(n, m)
	rect := make([][]bool, n)
	for i := 0; i < n; i++ {
		rect[i] = make([]bool, m)
	}

	isAvailable := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			return
		}
		if x >= n {
			ans = cnt
			return
		}
		// 检测下一行
		if y >= m {
			dfs(x+1, 0, cnt)
			return
		}
		// 如当前已经被覆盖，则直接尝试下一个位置
		if rect[x][y] {
			dfs(x, y+1, cnt)
			return
		}
		for k := common.Min(n-x, m-y); k >= 1 && isAvailable(x, y, k); k-- {
			// 将长度为 k 的正方形区域标记覆盖
			fillUp(x, y, k, true)
			// 跳过 k 个位置开始检测
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}
	}

	dfs(0, 0, 0)
	return ans
}
