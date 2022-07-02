package hard

import (
	"container/heap"
	"sort"
)

func minRefuelStops(target int, startFuel int, stations [][]int) (res int) {
	// 此处用一个优先队列保存每个油站的油，从大到小排列
	prev, h := 0, hp{}
	for i, allStation := 0, len(stations); i <= allStation; i++ {
		// 处理没有油站的边界情况
		cur := target
		if i < allStation {
			cur = stations[i][0]
		}
		// 计算跑到当前加油站消耗的油
		startFuel -= cur - prev
		// 如果油不够去下一站，就从队列里面加油
		for startFuel < 0 && h.Len() > 0 {
			startFuel += heap.Pop(&h).(int)
			res++
		}
		// 加了不够，就证明到不了，直接返回，这里也包含到第一站的边界，第一站没加都到不了就返回
		if startFuel < 0 {
			return -1
		}
		// 油够到下一站，就把当前行驶里程更新下，同时把当前油站到油放入优先队列
		if i < allStation {
			heap.Push(&h, stations[i][1])
			prev = cur
		}
	}
	// 循环完成时说明可以直接到目的地
	return res
}

type hp struct{ sort.IntSlice }

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
