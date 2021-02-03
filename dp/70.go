package dp

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	f1, result := 1, 2
	for i := 2; i < n; i++ {
		f1, result = result, result+f1
	}
	return result
}
