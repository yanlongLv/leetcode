package backTracking

func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backStaking func(cur int)
	backStaking = func(cur int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		if cur >= len(nums) {
			return
		}
		for i := cur; i < len(nums); i++ {
			path = append(path, nums[i])
			backStaking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backStaking(0)
	return result
}
