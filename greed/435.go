package greed

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	var res int
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			res++
			continue
		}
		end = intervals[i][1]
	}
	return res
}
