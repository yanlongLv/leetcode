package backTracking

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}
	sum := 0
	sort.Ints(candidates)
	var nowCombinate func(current int)
	nowCombinate = func(current int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := current; i < len(candidates) && sum+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			nowCombinate(i) // 这里不用i+1了表示此元素可以重复使用
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	nowCombinate(0)
	return result
}
