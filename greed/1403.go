package greed

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sums := 0
	for i := 0; i < len(nums); i++ {
		sums += nums[i]
	}
	subSums := 0
	result := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		subSums += nums[i]
		sums -= nums[i]
		result = append(result, nums[i])
		if subSums > sums {
			return result
		}
	}
	return result
}
