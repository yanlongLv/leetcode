package dp

func uniquePaths(m int, n int) int {
	var cache []int
	for i := 0; i < m; i++ {
		cache = append(cache, 1)
	}
	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			cache[i] += cache[i-1]
		}
	}
	return cache[m-1]
}
