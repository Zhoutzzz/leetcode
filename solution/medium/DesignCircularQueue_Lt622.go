package medium

type MyCircularQueue struct {
	head  int
	tail  int
	queue []int
	size  int
}

func Constructor(k int) MyCircularQueue {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularQueue{
		0,
		0,
		queue,
		0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.tail == len(this.queue) {
		this.tail = 0
	}
	this.queue[this.tail] = value
	this.tail++
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == len(this.queue) {
		this.head = 0
	}
	this.queue[this.head] = -1
	this.head++
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]

}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail-1]
}

func (this *MyCircularQueue) IsEmpty() (ans bool) {
	ans = this.size == 0
	if ans {
		this.tail = 0
		this.head = 0
	}
	return
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
