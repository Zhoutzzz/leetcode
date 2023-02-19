package medium

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	h := hpp(classes)
	heap.Init(&h)
	for ; extraStudents > 0; extraStudents-- {
		h[0][0]++
		h[0][1]++
		heap.Fix(&h, 0)
	}
	for _, c := range h {
		ans += float64(c[0]) / float64(c[1])
	}
	return ans / float64(len(classes))
}

type hpp [][]int

func (h hpp) Len() int { return len(h) }
func (h hpp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return (a[1]-a[0])*b[1]*(b[1]+1) > (b[1]-b[0])*a[1]*(a[1]+1)
}
func (h hpp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hpp) Push(interface{})     {}
func (hpp) Pop() (_ interface{}) { return }
