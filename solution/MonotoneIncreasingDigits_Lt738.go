package main

import "strconv"

func monotoneIncreasingDigits(N int) int {
	strArr := []byte(strconv.Itoa(N))
	i := 1
	for i < len(strArr) && strArr[i] > strArr[i-1] {
		i++
	}
	if i < len(strArr) {
		for i > 0 && strArr[i] < strArr[i-1] {
			i--
			strArr[i-1]--
		}
		for i++; i < len(strArr); i++ {
			strArr[i] = '9'
		}
	}
	res, _ := strconv.Atoi(string(strArr))
	return res
}
