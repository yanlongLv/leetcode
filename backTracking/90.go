package backTracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make(map[int]int)
	sort.Ints(nums)

	var backTracking func(current int)
	backTracking = func(current int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		if current > len(nums) {
			return
		}
		for i := current; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}
			used[i] = 1
			path = append(path, nums[i])
			backTracking(i + 1)
			used[i] = 0
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return result
}
