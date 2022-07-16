package easy

type MovingAverage struct {
	window []int
	size   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{make([]int, 0), size}
}

func (this *MovingAverage) Next(val int) (res float64) {
	this.window = append(this.window, val)
	if len(this.window) > this.size {
		this.window = this.window[1:]
	}

	for _, v := range this.window {
		res += float64(v)
	}
	return res / float64(len(this.window))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
