package main

func canCompleteCircuit(gas []int, cost []int) int {
	curGas := 0
	start := 0
	for ; start < len(gas); start++ {
		idx := start
		curGas = gas[start]
		for ; curGas-cost[idx] >= 0; curGas += gas[idx] {
			curGas = curGas - cost[idx]
			idx++
			if idx == len(gas) {
				idx = 0
			}
			if idx == start {
				return start
			}
		}
	}

	return -1
}
