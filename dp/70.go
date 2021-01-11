package dp

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f1, result := 1, 2
	for i := 2; i < n; i++ {
		f1, result = result, result+f1
	}
	return result
}
