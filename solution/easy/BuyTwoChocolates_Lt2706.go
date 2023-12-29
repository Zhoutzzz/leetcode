package easy

import "math"

func buyChoco(prices []int, money int) int {
	fi, se := math.MaxInt64, math.MaxInt64
	for _, price := range prices {
		if price < fi {
			se, fi = fi, price
		} else if price < se {
			se = price
		}
	}
	if money < fi+se {
		return money
	}
	return money - fi - se
}
