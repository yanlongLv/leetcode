package greed

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := 0
	for i := 0; i <= max; i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}
