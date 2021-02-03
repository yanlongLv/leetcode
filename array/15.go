package array

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	length := len(nums)
	if length < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}
		}
	}
	return result
}
