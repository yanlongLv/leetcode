package dp

func minCostClimbingStairs(cost []int) int {
	dp0, dp1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp0, dp1 = dp1, min(dp0, dp1)+cost[i]
	}
	return min(dp0, dp1)

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
