package backTracking

func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTracking func(current int)
	backTracking = func(current int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		if current > len(nums) {
			return
		}
		for i := current; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return result
}
