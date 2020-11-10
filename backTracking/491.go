package backTracking

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backTracking func(current int)
	backTracking = func(current int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
		if current > len(nums) {
			return
		}

		used := make(map[int]int) // 使用set对本层进行去重, 只记录本层
		for i := current; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]] == 1 {
				continue
			} else {
				used[nums[i]] = 1
				path = append(path, nums[i])
				backTracking(i + 1)
				path = path[:len(path)-1]
			}

		}
	}
	backTracking(0)
	return result
}
