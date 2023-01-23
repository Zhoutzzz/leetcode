package easy

import "leetcode/solution/common"

func calculateTax(brackets [][]int, income int) float64 {
	totalTax := 0
	lower := 0
	for _, bracket := range brackets {
		upper, percent := bracket[0], bracket[1]
		tax := (common.Min(income, upper) - lower) * percent
		totalTax += tax
		if income <= upper {
			break
		}
		lower = upper
	}
	return float64(totalTax) / 100
}
