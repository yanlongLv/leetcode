package array

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] == nums[i+int(len(nums)/2)] {
			return nums[i]
		}
	}
	return -1
}
