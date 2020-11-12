package array

import "math"

func minSubArrayLen(s int, nums []int) int {
	start := 0
	sum := 0
	min := math.MaxInt64
	subLength := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= s {
			subLength = end - start + 1
			sum -= nums[start]
			start++
			if subLength <= min {
				min = subLength
			}
		}
	}
	if min == math.MaxInt64 {
		min = 0
	}
	return min
}
