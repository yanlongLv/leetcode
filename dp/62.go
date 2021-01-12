package dp

// func uniquePaths(m int, n int) int {
// 	var cache []int
// 	for i := 0; i < m; i++ {
// 		cache = append(cache, 1)
// 	}
// 	for j := 1; j < n; j++ {
// 		for i := 1; i < m; i++ {
// 			cache[i] += cache[i-1]
// 		}
// 	}
// 	return cache[m-1]
// }

func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		cache[i][0] = 1
	}
	for j := 0; j < n; j++ {
		cache[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cache[i][j] = cache[i-1][j] + cache[i][j-1]
		}
	}
	return cache[m-1][n-1]
}
