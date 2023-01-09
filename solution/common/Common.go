package common

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func Gcd(a, b int) int {
	if b != 0 {
		return Gcd(b, a%b)
	}
	return a
}

func BinarySearch(nums []int, start, end, target int) int {
	mid := (start + end) / 2

	if nums[mid] > target {
		mid = BinarySearch(nums, start, mid-1, target)
	} else if nums[mid] < target {
		mid = BinarySearch(nums, mid, end, target)
	}
	return mid
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func QuickSort(arr [][]int, left, right int) [][]int {
	if left < right {
		partitionIndex := Partition(arr, left, right)
		QuickSort(arr, left, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func Partition(arr [][]int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i][0] < arr[pivot][0] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Equal(a, b []int) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
