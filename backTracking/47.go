package backTracking

import "sort"

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make(map[int]int)
	sort.Ints(nums)
	var backTracking func(current int)
	backTracking = func(current int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i] == used[i-1] && used[i-1] == 0 {
				continue
			}
			if used[i] == 0 {
				path = append(path, nums[i])
				used[i] = 1
				backTracking(i + 1)
				used[i] = 0
				path = path[:len(path)-1]
			}

		}
	}
	backTracking(0)
	return result
}
