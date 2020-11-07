package backTracking

func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTracking func(cur int)
	backTracking = func(cur int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// 剪枝操作 n - cur + len(path) + 1 > k
		if n-cur+len(path)+1 < k {
			return
		}
		for i := cur; i <= n; i++ {
			path = append(path, i)
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(1)
	return result
}
