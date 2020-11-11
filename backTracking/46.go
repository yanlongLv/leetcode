package backTracking

func permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make(map[int]int)
	var backTracking func(current int)
	backTracking = func(current int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}
			path = append(path, nums[i])
			used[i] = 1
			backTracking(i + 1)
			path = path[:len(path)-1]
			used[i] = 0
		}
	}
	backTracking(0)
	return result
}
