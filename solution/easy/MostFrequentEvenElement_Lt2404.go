package easy

func mostFrequentEven(nums []int) int {
	count := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 {
			count[x]++
		}
	}
	res, ct := -1, 0
	for k, v := range count {
		if res == -1 || ct < v || ct == v && k < res {
			res = k
			ct = v
		}
	}
	return res
}
