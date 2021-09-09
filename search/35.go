package search

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return right
}
