package backTracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}
	sum := 0
	sort.Ints(candidates)
	used := make(map[int]int)
	var backTracking func(current int)
	backTracking = func(current int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := current; i < len(candidates) && sum+candidates[i] <= target; i++ {
			// 去重复需要判断是指还是层
			// used[i - 1] == 1，说明同一树支candidates[i - 1]使用过
			// used[i - 1] == 0，说明同一树层candidates[i - 1]使用过
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == 0 {
				continue
			}
			used[i] = 1
			sum += candidates[i]
			path = append(path, candidates[i])
			backTracking(i + 1)
			used[i] = 0
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return result
}
