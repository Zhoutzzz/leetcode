package medium

func rearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) < 2 {
		return barcodes
	}

	counts := make(map[int]int)
	maxCount := 0
	for _, b := range barcodes {
		counts[b] = counts[b] + 1
		if counts[b] > maxCount {
			maxCount = counts[b]
		}
	}

	evenIndex := 0
	oddIndex := 1
	halfLength := len(barcodes) / 2
	res := make([]int, len(barcodes))
	for x, count := range counts {
		for count > 0 && count <= halfLength && oddIndex < len(barcodes) {
			res[oddIndex] = x
			count--
			oddIndex += 2
		}
		for count > 0 {
			res[evenIndex] = x
			count--
			evenIndex += 2
		}
	}
	return res
}
