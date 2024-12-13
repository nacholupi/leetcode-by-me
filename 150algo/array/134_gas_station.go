package array

func canCompleteCircuit(gas []int, cost []int) int {
	// 1. check if the total gas is enough to travel the whole circuit
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}

	// 2. find the starting point
	// if the total gas is enough, there must be a starting point
	// that can travel the whole circuit
	start, tank := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	return start
}
