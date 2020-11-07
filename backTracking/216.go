package backTracking

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	path := []int{}
	sum := 0
	var backstacking func(cur int)
	backstacking = func(cur int) {
		if sum == n && len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		if sum > n || len(path) > k {
			return
		}
		for i := cur; i <= 9; i++ {
			path = append(path, i)
			sum += i
			backstacking(i + 1)
			path = path[:len(path)-1]
			sum -= i
		}
	}
	backstacking(1)
	return result
}
