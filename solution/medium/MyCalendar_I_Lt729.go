package medium

type MyCalendar struct {
	calender [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{calender: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	calenderLen := len(this.calender)
	flag := true
	for i := 0; i < calenderLen; i++ {
		if calenderLen == 0 || start >= this.calender[i][1] || end <= this.calender[i][0] {
			flag = true
		} else {
			return false
		}
	}
	if flag {
		this.calender = append(this.calender, []int{start, end})
	}
	return flag
}

/*

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)


二分搜索
type MyCalendar struct {
    *redblacktree.Tree
}

func Constructor() MyCalendar {
    t := redblacktree.NewWithIntComparator()
    t.Put(math.MaxInt32, nil) // 哨兵，简化代码
    return MyCalendar{t}
}

func (c MyCalendar) Book(start, end int) bool {
    node, _ := c.Ceiling(end)
    it := c.IteratorAt(node)
    if !it.Prev() || it.Value().(int) <= start {
        c.Put(start, end)
        return true
    }
    return false
}
*/
