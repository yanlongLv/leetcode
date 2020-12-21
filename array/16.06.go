package array

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	result := 0
	i, j, result := 0, 0, math.MaxInt32
	for i < len(a) && j < len(b) {
		result = min(result, abs(a[i], b[j]))
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	return result
}

func abs(x, y int) int {
	if x-y < 0 {
		return y - x
	}
	return x - y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
