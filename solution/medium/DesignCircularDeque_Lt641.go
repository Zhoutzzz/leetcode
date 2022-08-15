package medium

type MyCircularDeque struct {
	Head     *QueueNode
	Tail     *QueueNode
	Capacity int
	Size     int
}

type QueueNode struct {
	Val  int
	Prev *QueueNode
	Next *QueueNode
}

func Constructor6(k int) MyCircularDeque {
	deque := MyCircularDeque{
		nil, nil, k, 0,
	}
	return deque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Size++
	if this.Head == nil {
		this.Head = &QueueNode{
			Val:  value,
			Prev: nil,
			Next: nil,
		}
		this.Tail = this.Head
		return true
	}
	prev := this.Head
	this.Head.Prev = &QueueNode{
		Val:  value,
		Prev: nil,
		Next: nil,
	}
	this.Head = this.Head.Prev
	this.Head.Next = prev
	this.Head.Prev = this.Tail
	this.Tail.Next = this.Head

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Size++
	if this.Tail == nil {
		this.Tail = &QueueNode{
			Val:  value,
			Prev: nil,
			Next: nil,
		}
		this.Head = this.Tail
		return true
	}
	prev := this.Tail
	this.Tail.Next = &QueueNode{
		Val:  value,
		Prev: nil,
		Next: nil,
	}
	this.Tail = this.Tail.Next
	this.Tail.Prev = prev
	this.Head.Prev = this.Tail
	this.Tail.Next = this.Head

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Size--
	if this.Head == this.Tail {
		this.Head = nil
		this.Tail = nil
		return true
	}
	this.Head = this.Head.Next
	this.Tail.Next = this.Head
	this.Head.Prev = this.Tail
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Size--
	if this.Head == this.Tail {
		this.Head = nil
		this.Tail = nil
		return true
	}
	this.Tail = this.Tail.Prev
	this.Tail.Next = this.Head
	this.Head.Prev = this.Tail
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
