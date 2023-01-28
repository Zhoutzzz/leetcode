package medium

func waysToMakeFair(nums []int) (res int) {
	var odd1, even1, odd2, even2 int
	for i, num := range nums {
		if i&1 > 0 {
			odd2 += num
		} else {
			even2 += num
		}
	}
	for i, num := range nums {
		if i&1 > 0 {
			odd2 -= num
		} else {
			even2 -= num
		}
		if odd1+even2 == odd2+even1 {
			res++
		}
		if i&1 > 0 {
			odd1 += num
		} else {
			even1 += num
		}
	}
	return
}
