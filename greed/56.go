package greed

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else if intervals[i][1] > result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result
}
