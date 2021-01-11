package dp

func fib(n int) int {
	f0, result := 0, 1
	if n <= 0 {
		return f0
	}
	for i := 1; i < n; i++ {
		result, f0 = result+f0, result
	}
	return result
}
