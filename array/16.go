package array

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 1<<31 - 1
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			add := nums[i] + nums[left] + nums[right]
			if add == target {
				return add
			}
			if abs(target-add) < abs(target-result) {
				result = add
			}
			if target > add {
				left += 1
			}
			if target < add {
				right -= 1
			}
		}

	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
