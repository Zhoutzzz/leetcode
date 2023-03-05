package medium

import "leetcode/solution/common"

func minOperationsMaxProfit(customers []int, boardingCost, runningCost int) int {
	ans := -1
	var maxProfit, totalProfit, operations, customersCount int
	for _, c := range customers {
		operations++
		customersCount += c
		curCustomers := common.Min(customersCount, 4)
		customersCount -= curCustomers
		totalProfit += boardingCost*curCustomers - runningCost
		if totalProfit > maxProfit {
			maxProfit = totalProfit
			ans = operations
		}
	}
	if customersCount == 0 {
		return ans
	}
	profitEachTime := boardingCost*4 - runningCost
	if profitEachTime <= 0 {
		return ans
	}
	if customersCount > 0 {
		fullTimes := customersCount / 4
		totalProfit += profitEachTime * fullTimes
		operations += fullTimes
		if totalProfit > maxProfit {
			maxProfit = totalProfit
			ans = operations
		}
		remainingCustomers := customersCount % 4
		remainingProfit := boardingCost*remainingCustomers - runningCost
		totalProfit += remainingProfit
		if totalProfit > maxProfit {
			maxProfit = totalProfit
			operations++
			ans++
		}
	}
	return ans
}
