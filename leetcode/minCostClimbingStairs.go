package leetcode

func MinCostClimbingStairs(cost []int) int {
	return minCostClimbingStairs(cost)
}

func minCostClimbingStairs(cost []int) int {
	stepState := make([]int, len(cost))
	skipState := make([]int, len(cost))

	stepState[0] = cost[0]
	skipState[0] = 0

	for i := 1 ; i < len(cost) ; i++ {
		stepState[i] = minCostStep(stepState[i-1], skipState[i-1]) + cost[i]
		skipState[i] = stepState[i-1]
	}

	return minCostStep(stepState[len(cost)-1], skipState[len(cost)-1])
}

func minCostStep(a, b int) int {
	if a < b {
		return a
	}
	return b
}